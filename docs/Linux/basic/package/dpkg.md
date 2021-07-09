# dpkg

package manager for Debian

## dpkg

- `-i | --install <package...>`

  安装pkg

- `-r | --remove <package...>`

  删除pkg，但是不会将配置文件删除

- `-P | --purge <package...>`

  删除pkg同时删除配置文件，和`apt-get purge`相同

- `-V | --verify [package]`

  校验pkg的完整性，如果忽略package默认校验所有的pkg

- `--add-architecture <architecture>`

  安装不同架构的dpkg时可以直接安装不使用`--force-architecture`。可以使用`--print-architecture`可以输出当前使用architecture

## dkpg-query

- -l | --list

  查看所有安装的pkg，可以配合`grep`一起

- `-S | --search <filename>`

  从安装的pkg中所有指定file

  ```
  dpkg -S kvm | head -10
  ```

  可以通过该命令过滤出某个命令归属的包
  
  ```
  root in /home/ubuntu λ dpkg-query --search $(which arp)
  net-tools: /usr/sbin/arp
  ```
  
  

