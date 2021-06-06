# kubectl apply

将配置(file ro  stdin)应用于一个资源，如果这个资源不存在会被自动创建。接受JSON和YAML格式。

syntax：`kubectl apply (-f filename | -k directory) `

## 参数

- -f | --filename

  应用的配置文件

- --kustomize | -k

  指定一个目录包含配置文件

  ```
  [root@k8smaster opt]# kubectl apply -k /opt
  error: unable to find one of 'kustomization.yaml', 'kustomization.yml' or 'Kustomization' in directory '/opt'
  ```

- --dry-run

- --output | -o

