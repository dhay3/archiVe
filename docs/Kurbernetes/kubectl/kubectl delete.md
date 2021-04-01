# kubectl delete

删除kubernetes resources

syntax：`kubectl delete ([-f filename] | [-k direcotory] | [type])`

```
kubectl delete -f ./pod.json                                              # Delete a pod using the type and name specified in pod.json
kubectl delete pod,service baz foo                                        # Delete pods and services with same names "baz" and "foo"
kubectl delete pods,services -l name=myLabel                              # Delete pods and services with label name=myLabel
kubectl -n my-ns delete pod,svc --all                                      # Delete all pods and services in namespace my-ns,
# Delete all pods matching the awk pattern1 or pattern2
kubectl get pods  -n mynamespace --no-headers=true | awk '/pattern1|pattern2/{print $1}' | xargs  kubectl delete -n mynamespace pod
```

## 问题解决

如果pod通过workload创建的，使用`kubectl delete pods --all`无法删除pods，==需要删除workload==

```
[root@k8smaster opt]# kubectl delete deployments --all
```

