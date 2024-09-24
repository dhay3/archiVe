# Union all 替代

有些系统会自动屏蔽 union 关键字，这里巧妙的使用了一个新增字段 flag 并赋值不同，然后关联取非空值。可以绕开明文屏蔽 union 的策略

```
select
  snat_seg,
  snat_type,
  cluster
from(
    select
      case
        when an_seg is null then sn_seg
        else an_seg
      end as snat_seg,
      case
        when an_type is null then sn_type
        else an_type
      end as snat_type,
      case
        when an_cluster is null then sn_cluster
        else an_cluster
      end as cluster
    from
      (
        select
          snat_seg as an_seg,
          snat_type as an_type,
          cluster as an_cluster,
          1 as flag
        from
          an_snat_segments
      ) as an full
      join (
        select
          snat_seg as sn_seg,
          snat_type as sn_type,
          cluster as sn_cluster,
          2 as flag
        from
          sn_snat_segments
      ) as sn on an.flag = sn.flag
  ) as t
order by
  cluster,
  snat_type,
  snat_seg;
```

