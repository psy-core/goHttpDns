#!/usr/bin/env bash

# 进入脚本所在目录
cd `dirname $0`
rm -rf ./dist
mkdir -p dist/bin
cp -r ./etc dist

#go build
echo go build...
cd main
go build
mv main ../dist/bin/goHttpDNS

# generate start.sh
echo generate start script...
cd ../dist/bin
echo '#!/usr/bin/env bash' > start.sh
echo >> start.sh
echo 'cd `dirname $0`' >> start.sh
echo 'EXEDIR=`pwd`' >> start.sh
echo >> start.sh
echo '$EXEDIR/goHttpDNS' >> start.sh
chmod a+x start.sh

echo done.

# 导出环境变量
# echo "[packing] exporting gopath..."
# export GOPATH=`pwd`

# echo "[packing] cd src/main..."
# cd src/main

#go install
# echo "[packing] [go installing...]"
# go install -ldflags "-w"
# if [[ $? -ne 0 ]] ; then
#     echo "[packing] [go install error]"
#     echo "[packing] [pack failed]"
#     exit 1
# fi
# echo "[packing] [go install success]"

#packing
# echo "[packing] cd ../../bin ..."
# cd ../../bin
# echo "[packing] renameing to emoticonservice..."
# mv main emoticonservice
# echo "[packing] [pack success]"
