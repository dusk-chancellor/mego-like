i=60
while [[ $(kubectl get pods -n like --no-headers=true|grep like|awk '{print $2}'|uniq) != "1/1" ]]
  do
 if [ $i == 0 ]; then 
  echo "No more retries left. Deployment failed."
  exit 1
 fi 
 ((i--))
 echo "-----------------------"
 echo "Retries left: $i"
 kubectl get pods -n like
 sleep 10
done