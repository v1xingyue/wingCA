 1  shutdown -h now
    2  shutdown -h now
    3  ls
    4  ps aux | grep wildcat
    5  /usr/local/wildcat/wildcat -v
    6  /usr/local/wildcat/wildcat -v
    7  ls
    8  ls
    9  ps aux 
   10  openssl 
   11  ls
   12  ymkdir ca && cd ca  
   13  mkdir ca && cd ca  
   14  mkdir newcerts private conf server
   15  ls
   16  lls
   17  ls
   18  vim conf/openssl.conf
   19  ls
   20  openssl genrsa -out private/ca.key 2048  
   21  openssl req -new -key private/ca.key -out private/ca.csr  
   22  openssl x509 -req -days 365 -in private/ca.csr -signkey private/ca.key -out private/ca.crt  
   23  echo FACE > serial  
   24  #可以是任意四个字符
   25  touch index.txt  
   26  openssl ca -gencrl -out ./private/ca.crl -crldays 7 -config "./conf/openssl.conf"  
   27  openssl genrsa -out server/server.key 2048  
   28  openssl req -new -key server/server.key -out server/server.csr  
   29  openssl ca -in server/server.csr -cert private/ca.crt -keyfile private/ca.key -out server/server.crt -config "./conf/openssl.conf"  
   30  vim conf/openssl.conf 
   31  openssl ca -in server/server.csr -cert private/ca.crt -keyfile private/ca.key -out server/server.crt -config "./conf/openssl.conf"  
   32  ls
   33  mkdir users  
   34  openssl genrsa -des3 -out ./users/client.key 2048  
   35  ls
   36  ls users/client.key 
   37  openssl req -new -key ./users/client.key -out ./users/client.csr  
   38  openssl ca -in ./users/client.csr -cert ./private/ca.crt -keyfile ./private/ca.key -out ./users/client.crt -config "./conf/openssl.conf"  
   39  openssl pkcs12 -export -clcerts -in ./users/client.crt -inkey ./users/client.key -out ./users/client.p12  
   40  openssl pkcs12 -export -clcerts -in ./users/client.crt -inkey ./users/client.key -out ./users/client.p12  
   41  ls users/client.p12 
   42  ls
   43  yum install nginx
   44  cd /etc/nginx/conf.d/
   45  ls
   46  ls
   47  cd ..
   48  ls
   49  vim nginx.conf
   50  pwd
   51  ls /usr/home/xingyue/ca/server/server.key 
   52  vim nginx.conf
   53  systemctl restart nginx
   54  nginx
   55  vim nginx.conf
   56  nginx
   57  ls "/usr/home/xingyue/ca/server/server.key" 
   58  ls /usr/home/xingyue/ca/server/server.key 
   59  cat /usr/home/xingyue/ca/server/server.key
   60  nginx
   61  vim nginx.conf
   62  nginx
   63  ifconfig 
   64  ls
   65  cd 
   66  ls
   67  cd /usr/home/xingyue/cli
   68  cd /usr/home/xingyue/ca/cli
   69  ls
   70  cd /usr/home/xingyue/ca/
   71  ls
   72  cd users/
   73  ls
   74  ls
   75  ll
   76  ll
   77  cat client.p12 
   78  ls
   79  ls
   80  ll
   81  python -m SimpleHTTPServer
   82  ls
   83  cd 
   84  ls
   85  exit
   86  ls
   87  ls
   88  cd ca/
   89  ls
   90  ll
   91  cd users/
   92  ls
   93  cd ..
   94  ls
   95  openssl req -new -key ./users/client.key -out ./users/client.csr  
   96  ls
   97  ll users/
   98  openssl ca -in ./users/client.csr -cert ./private/ca.crt -keyfile ./private/ca.key -out ./users/client.crt -config "./conf/openssl.conf"  
   99  ll users/
  100  openssl pkcs12 -export -clcerts -in ./users/client.crt -inkey ./users/client.key -out ./users/client.p12  
  101  ls
  102  cd users/
  103  python -m SimpleHTTPServer
  104  ls
  105  cd ..
  106  ls
  107  openssl req -new -key ./users/client.key -out ./users/client.csr  
  108  openssl ca -in ./users/client.csr -cert ./private/ca.crt -keyfile ./private/ca.key -out ./users/client.crt -config "./conf/openssl.conf"  
  109  openssl pkcs12 -export -clcerts -in ./users/client.crt -inkey ./users/client.key -out ./users/client.p12  
  110  ls
  111  ls
  112  cd users/
  113  ls
  114  python -m SimpleHTTPServer
  115  ls
  116  ls
  117  ls
  118  yum install mongo
  119  vim /etc/yum.repos.d/mongodb.repo
  120  yum install -y mongodb-org
  121  systemctl start mongodb
  122  ls
  123  yum install -y mongodb-org
  124  systemctl start mongodb-org
  125  systemctl start mongo
  126  /etc/init.d/mongod start 
  127  ls
  128  vim /etc/mongod.conf 
  129  /etc/init.d/mongod restart 
  130  netstat -antuple | grep 27017
  131  ifconfig 
  132  mongo
  133  vim /etc/mongod.conf 
  134  /etc/init.d/mongod restart 
  135  netstat -antuple | grep 278
  136  netstat -antuple | grep 27
  137  vim /etc/mongod.conf 
  138  /etc/init.d/mongod restart 
  139  ifconfig | grep 27
  140  netstat -antuple | grep 27
  141  netstat -antuple | grep 27
  142  /etc/init.d/mongod restart 
  143  vim /etc/mongod.conf 
  144  ls
  145  /etc/init.d/mongod restart 
  146  vim /etc/mongod.conf 
  147  ls
  148  pstree
  149  ls
  150  ps aux 
  151  yum install psutils-y 
  152  yum install psutils-y 
  153  yum install psutils -y 
  154  pstre
  155  pstree
  156  ls
  157  yum provides pstree
  158  yum install psmisc -y 
  159  ls
  160  ls
  161  ll
  162  ls
  163  pstree
  164  ps aux | grep mongo
  165  ps aux | grep mongo
  166  vim /etc/mongod.conf 
  167  ls
  168  /usr/local/wildcat/wildcat -v
  169  /usr/local/wildcat/wildcat -v
  170  /usr/local/wildcat/wildcat -v
  171  ls
  172  ps aux | grep cattag
  173  ps aux | grep cat
  174  ls
  175  ifconfig 
  176  ls
  177  ll
  178  whoami 
  179  top -d 1
  180  ls
  181  ifconfig 
  182  ls
  183  ifconfig 
  184  ls
  185  ll
  186  pstree
  187  ls
  188  netstat -antuple  | grep ngin x
  189  netstat -antuple  | grep nginx
  190  cd /etc/nginx/
  191  ls
  192  cd conf.d/
  193  ls
  194  ll
  195  cd ..
  196  ls
  197  vim nginx.conf
  198  cd /usr/home/xingyue/
  199  ls
  200  ll
  201  cd ca/
  202  ls
  203  ll
  204  ls
  205  cd users/
  206  ls
  207  ll
  208  cd ..
  209  ls
  210  ifconfig 
  211  ls
  212  cd /etc/nginx/conf.d/
  213  ls
  214  cd ..
  215  ls
  216  cd -
  217  ls
  218  cd /usr/home/xingyue/
  219  ls
  220  cd ca/
  221  ls
  222  git remote -v 
  223  ls
  224  ll
  225  vim index.txt
  226  ls
  227  cd conf/
  228  ls
  229  ll
  230  vim openssl.conf 
  231  ls
  232  cd ..
  233  ls
  234  cd newcerts/
  235  ls
  236  ll
  237  cd ..
  238  ls
  239  cd ..
  240  ls
  241  lsls
  242  ll
  243  ls
  244  ifconfig 
  245  ls
  246  cd ca/
  247  ls
  248  ll
  249  ls
  250  find . -name shell
  251  cd /root/
  252  ls
  253  ll
  254  cd -
  255  ls
  256  ll
  257  cd private/
  258  ls
  259  ll
  260  ls
  261  ll
  262  ls
  263  ll
  264  cd ..
  265  ls
  266  history 
