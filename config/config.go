package config

const yamlFile = `

development:
  port: :9091 
  database: txhims_staging
  username: root
  password: n[h^nM06
  host: test.tongxinyiliao.com
  dbport: 3306
  redis_host: '127.0.0.1'
  redis_port: 6379
  redis_pwd: ''
  redis_db: 3

staging:
  port: :9091
  database: txhims_staging
  username: root
  password: n[h^nM06
  host: test.tongxinyiliao.com
  dbport: 3306
  redis_host: 'dicomuptest.tongxinyiliao.com'
  redis_port: 6379
  redis_pwd: 'b1e41f0f3dc551703dc51c64d1970fe5'
  redis_db: 3

production:
  port: :9091
  database: txhims_production
  username: pro_txdb
  password: Tongxin2014
  host: rdsarjziaiq3e2a.mysql.rds.aliyuncs.com
  dbport: 3306
  redis_host: 'r-2ze1efa5198c21c4.redis.rds.aliyuncs.com'
  redis_port: 6379
  redis_pwd: 'r-2ze1efa5198c21c4:TongXin2017'
  redis_db: 3




`
