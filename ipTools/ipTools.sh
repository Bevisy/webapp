#!/usr/bin/env sh

set -x
#KONGVIP=10.125.29.13/22

# return 0(valid), 1(not valid)
#function isValidIp() {
#  local ip=$1
#  local ret=1
#
#  if [[ $ip =~ ^[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}$ ]]; then
#    ip=(${ip//\./ }) # 按.分割，转成数组，方便下面的判断
#    [[ ${ip[0]} -le 255 && ${ip[1]} -le 255 && ${ip[2]} -le 255 && ${ip[3]} -le 255 ]]
#    ret=$?
#  fi
#
#  return $ret
#}

# KONGVIP 值判断
if [[ ${KONGVIP} ]];then
  echo "KONG VIP is ${KONGVIP}"
  VIP=`echo ${KONGVIP} | awk -F'/' '{print $1}'`
  PREFIX=`echo ${KONGVIP} | awk -F'/' '{print $2}'`
#  if ! isValidIp ${VIP};then
#    echo "$VIP IS NOT VALID!"
#    exit 1
#  else
#    echo "${VIP} IS VALID."
#  fi
else
  echo "KONG VIP IS NOT EXIST!"
  exit 1
fi

# 判断传入的IP是否已存在
#return 1(failure) not used,0(success) used
function isUsed {
  if ping -c 2 $1 >/dev/null; then
      return 0
  else
      return 1
  fi
}

# 判断需要添加的网卡名称

function ipNetworkAddress() {
    ipcalc -n $1 | awk -F'=' '{print $2}'
}

# inet6\|docker0\|tun0\|127.0.0.1\|secondary
# 排除Ipv6、docker、tun、lo、网卡多绑IP情况
for i in `ip addr | grep inet | grep -v 'inet6\|docker0\|tun0\|127.0.0.1\|secondary' | awk '{print $2"|"$NF}'`
  do
    if [[ "$(ipNetworkAddress $(echo $i | awk -F'|' '{print $1}'))" = "$(ipNetworkAddress ${KONGVIP})" ]];then
      TargetInterface=$(echo $i | awk -F'|' '{print $2}')
    fi
done

# 添加网卡IP
if [[ "$1" = "add" ]];then
  # IP 不存在则添加， 存在则报错退出
  if isUsed ${VIP};then
    echo "$VIP IS IN USED, PLEASE CHANGE ANOTHER ONE!"
    exit 1
  else
    echo "$VIP IS NOT USED."
    ip addr add ${KONGVIP} dev ${TargetInterface}
  fi
elif [[ "$1" = "delete" ]];then
  # IP 存在则删除，不存在则正常退出
  if isUsed ${VIP};then
    echo "$VIP IS IN USED, DELETE it."
    ip addr del ${KONGVIP} dev ${TargetInterface}
  else
    echo "$VIP IS NOT USED."
  fi
fi
