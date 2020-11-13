```
kubectl run mysql-client --image=mysql:5.7 -i --rm --restart=Never --\
  mysql -h mysql-0.mysql -ppassword <<EOF
CREATE DATABASE todo_app;
EOF
```