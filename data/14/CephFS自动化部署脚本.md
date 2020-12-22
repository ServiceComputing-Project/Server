## 前提
1. 主机为CentOS 7.2及以上
2. 预装ceph版本为ceph-luminous 12.2.5
3. 主节点可无密访问其他两个节点
4. 搭建3台主机形成的ceph cluster及cephfs
5. 主节点配置本地ceph yum源

配置详情如下：

主节点: ceph-deploy & mon & mgr & osd & ntp &mds &httpd & createrepo

从节点: mon & mgr & osd & ntp &mds

从节点: mon & mgr & osd & ntp &mds

## 目录结构

```
[root@ceph-1 deploy]# ll
total 132
-rw-r--r-- 1 root root 97527 Jun 21 17:09 deploy.log
-rwxr-xr-x 1 root root  4288 Jun 21 17:02 deploy.sh
drwxr-xr-x 2 root root  4096 Jun 21 16:54 ntp
-rwxr-xr-x 1 root root    99 Jun 21 10:43 part.sh
drwxr-xr-x 2 root root  4096 Jun 20 15:57 py
drwxr-xr-x 2 root root  4096 Jun 20 10:46 rpm
drwxr-xr-x 2 root root  4096 Jun 21 16:39 test
drwxr-xr-x 2 root root  4096 Jun 21 16:54 yum
```
1. rpm存放的是ceph相关的rpm包，用于建立本地源，可参见搭建ceph本地源下载相关的rpm
2. ntp存放主节点的ntp.conf和从节点的ntp1.conf

ntp.conf如下

```
driftfile /var/lib/ntp/drift
restrict default nomodify
restrict 127.0.0.1 
restrict ::1

# Enable public key cryptography.
# #crypto
#
includefile /etc/ntp/crypto/pw
#
# # Key file containing the keys and key identifiers used when operating
# # with symmetric key cryptography. 
keys /etc/ntp/keys
#
server 127.127.1.0
fudge 127.127.1.0 stratum 10
```
ntp1.conf
```
driftfile /var/lib/ntp/drift
restrict default nomodify
restrict 127.0.0.1 
restrict ::1

# Enable public key cryptography.
#crypto

includefile /etc/ntp/crypto/pw

# Key file containing the keys and key identifiers used when operating
# with symmetric key cryptography. 
keys /etc/ntp/keys
```
3. yum中包含CentOS-Base.repo，epel.repo和ceph.repo

CentOS-Base.repo，epel.repo可参见Ceph 部署（Centos7 + Luminous）

ceph.repo
```
[ceph]
name=ceph
```
## 脚本 deploy.sh
```

sed -i 's/SELINUX=.*/SELINUX=disabled/' /etc/selinux/config
setenforce 0
ssh root@$ceph_2_name "sed -i 's/SELINUX=.*/SELINUX=disabled/' /etc/selinux/config;setenforce 0"
ssh root@$ceph_3_name "sed -i 's/SELINUX=.*/SELINUX=disabled/' /etc/selinux/config;setenforce 0"

#搭建本地yum源头
yum install httpd createrepo -y
mkdir -p /var/www/html/ceph/12.2.5
cp $DEPLOY_DIR/rpm/* /var/www/html/ceph/12.2.5
createrepo /var/www/html/ceph/12.2.5

#配置yum
yum clean all
echo "baseurl=http://$ceph_1_ip/ceph/12.2.5" >> $DEPLOY_DIR/yum/ceph.repo
echo "gpgcheck=0" >> $DEPLOY_DIR/yum/ceph.repo
\cp -fr $DEPLOY_DIR/yum/* /etc/yum.repos.d/
yum makecache

ssh root@$ceph_2_name "yum clean all"
ssh root@$ceph_3_name "yum clean all"
scp -r $DEPLOY_DIR/yum/* root@$ceph_2_name:/etc/yum.repos.d/
scp -r $DEPLOY_DIR/yum/* root@$ceph_3_name:/etc/yum.repos.d/ 
ssh root@$ceph_2_name "yum makecache"
ssh root@$ceph_3_name "yum makecache"

#安装ntp
yum install ntp -y
\cp -fr $DEPLOY_DIR/ntp/ntp.conf /etc/
echo "server $ceph_1_ip" >> $DEPLOY_DIR/ntp/ntp1.conf
systemctl enable ntpd
systemctl restart ntpd

ssh root@$ceph_2_name "yum install ntp -y"
ssh root@$ceph_3_name "yum install ntp -y"
scp -r $DEPLOY_DIR/ntp/ntp1.conf root@$ceph_2_name:/etc/ntp.conf
scp -r $DEPLOY_DIR/ntp/ntp1.conf root@$ceph_3_name:/etc/ntp.conf
ssh root@$ceph_2_name "systemctl enable ntpd;systemctl restart ntpd"
ssh root@$ceph_3_name "systemctl enable ntpd;systemctl restart ntpd"

#安装ceph
yum install ceph -y
ssh root@$ceph_2_name "yum install ceph -y"
ssh root@$ceph_3_name "yum install ceph -y"

#安装ceph-deploy
yum install ceph-deploy -y

#部署ceph
mkdir ~/cluster
cd ~/cluster
CLUSTER_DIR=$(cd `dirname $0`; pwd)
ceph-deploy new $ceph_1_name $ceph_2_name $ceph_3_name
echo "public_network=$sub_network" >> ceph.conf
#echo "osd_crush_update_on_start = false" >> ceph.conf

ceph-deploy mon create-initial
ceph-deploy admin $ceph_1_name $ceph_2_name $ceph_3_name

#配置osd
index=0
for dev_name in ${ceph_1_dev[@]}
do
ceph-volume lvm zap /dev/$dev_name
ceph-deploy osd create $ceph_1_name --bluestore --data /dev/$dev_name --block-db /dev/${ceph_1_dev_journal[$index]} --block-wal /dev/${ceph_1_dev_journal[$index+1]}
index=$[$index+2]
done
index=0
for dev_name in ${ceph_2_dev[@]}
do
ssh root@$ceph_2_name "ceph-volume lvm zap /dev/$dev_name"
ceph-deploy osd create $ceph_2_name --bluestore --data /dev/$dev_name --block-db /dev/${ceph_2_dev_journal[$index]} --block-wal /dev/${ceph_2_dev_journal[$index+1]}
index=$[$index+2]
done
index=0
for dev_name in ${ceph_3_dev[@]}
do
ssh root@$ceph_3_name "ceph-volume lvm zap /dev/$dev_name"
ceph-deploy osd create $ceph_3_name --bluestore --data /dev/$dev_name --block-db /dev/${ceph_3_dev_journal[$index]} --block-wal /dev/${ceph_3_dev_journal[$index+1]}
index=$[$index+2]
done


#配置mgr
ceph-deploy mgr create $ceph_1_name $ceph_2_name $ceph_3_name

#启动dashboard
ceph mgr module enable dashboard

#部署cephfs
ceph-deploy mds create $ceph_1_name $ceph_2_name $ceph_3_name
ceph osd pool create cephfs_data 128
ceph osd pool create cephfs_metadata 128
ceph fs new cephfs cephfs_metadata cephfs_data

mkdir /mnt/mycephfs  
admin_key=`ceph auth get-key client.admin`
admin_key_base64=`ceph auth get-key client.admin |base64`
sleep 5#等待mds部署完成后再mount
mount -t ceph $ceph_1_name:6789,$ceph_2_name:6789,$ceph_3_name:6789:/ /mnt/mycephfs -o name=admin,secret=$admin_key
```
## 执行脚本
```
nohup ./deploy.sh > deploy.log 2>&1 &
```

#### 本文连接： https://www.jianshu.com/p/5ca7990db0f1
