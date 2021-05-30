# packer builder null

null builder并不生成artifacts，一般只用作对provisioners进行校验，而不用等待build的过程。

```
{
  "builders": [
    {
      "type": "null"
    }
  ]
}
```

