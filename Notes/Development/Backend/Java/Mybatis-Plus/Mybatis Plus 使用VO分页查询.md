# Mybatis Plus 使用VO分页查询

[TOC]

现在想要分页查询数据库, 将查询出来的数据封装到`vo`, 如果在使用`limit`, 将无法通过一次查询获取总数据条数, 但是mp封装了

### #VO

这里组合了`course`, `subject`, `teacher`作为`DO`的实体类, 也是`sql`返回的`resultType`

```java
package com.chz.eduservice.entity.vo;

import com.chz.utils.statuscode.CourseStatus;
import io.swagger.annotations.ApiModel;
import lombok.Data;

import java.io.Serializable;

/**
 * @Author: CHZ
 * @DateTime: 2020/7/3 16:24
 * @Description: TODO
 */
@ApiModel("课程发布信息封装类,用于展示course和在course list中显示")
@Data
public class CoursePublishInfoVo implements Serializable {
    private static final long serialVersionUID = 1L;
    private String id;
    private String title;
    private String cover;
    private Integer lessonNum;
    private String subjectLevelOne;
    private String subjectLevelTwo;
    private String teacherName;
    private String price;
    private String viewCount;
    private CourseStatus status;
}

```

### #DAO

注意这里必须要有`Page`做参数, 泛型为想要作为分页查询一条记录的`VO`

```java
/**
     * 按条件分页查询course
     *
     * @param page 注意这里必须要有Page对象,否则mp无法完成分页查询
     * @param courseQuery
     * @return
     */
    List<CoursePublishInfoVo> pageCourseAllInfo(Page<CoursePublishInfoVo> page,
                                                @Param("courseQuery") CourseQuery courseQuery);

```

### #sql

这里一定不能使用`limit`

```sql
    <select id="pageCourseAllInfo" resultType="com.chz.eduservice.entity.vo.CoursePublishInfoVo">
        SELECT ec.id, ec.title, ec.price, ec.status, ec.lesson_num AS lessonNum,ec.view_count AS viewCount,
        et.name AS teacherName,
        es1.title AS subjectLevelOne,
        es2.title AS subjectLevelTwo
        FROM edu_course ec
        LEFT JOIN edu_course_description ecd ON ec.id = ecd.id
        LEFT JOIN edu_teacher et ON ec.teacher_id = et.id
        LEFT JOIN edu_subject es1 ON ec.subject_parent_id = es1.id
        LEFT JOIN edu_subject es2 ON ec.subject_id = es2.id
        <where>
            <if test="courseQuery.title!=null">
                AND ec.title = #{courseQuery.title}
            </if>
            <if test="courseQuery.teacherId!=null">
                AND et.id = #{courseQuery.teacherId}
            </if>
            <if test="courseQuery.subjectParentId!=null">
                AND ec.subject_parent_id = #{courseQuery.subjectParentId}
            </if>
            <if test="courseQuery.subjectId!=null">
                AND ec.subject_id = #{courseQuery.subjectId}
            </if>
            <if test="courseQuery.status!=null">
                AND ec.status = #{courseQuery.status}
            </if>
            <if test="courseQuery.beginPrice!=null and courseQuery.endPrice!=null">
                AND ec.price BETWEEN courseQuery.beginPrice AND courseQuery.endPrice
            </if>
        </where>
        ORDER BY ec.gmt_create DESC
    </select>

```

### #Service

```java
	@Override
    public Page<CoursePublishInfoVo> pageCourseAllInfo(Integer cur, Integer size, CourseQuery courseQuery) {
        Page<CoursePublishInfoVo> page = new Page<>(cur, size);
        //将查询结果封装到page中,作为page中的数据
        page.setRecords(baseMapper.pageCourseAllInfo(page, courseQuery));
        return page;
    }
```

### #Controller

```java
	@ApiOperation(value = "分页查询")
    @PostMapping("/{curPage}/{pageSize}")
    public ResponseBo pageCourseOnCondition(@Min(1) @PathVariable Integer curPage,
                                            @PathVariable Integer pageSize,
                                            @RequestBody CourseQuery courseQuery) {
        HashMap<String, Object> map = new HashMap<>();
        Page<CoursePublishInfoVo> page = courseService.pageCourseAllInfo(curPage, pageSize, courseQuery);
        map.put("courses",page.getRecords());
        map.put("total",page.getTotal());
        return ResponseBo.ok().data("info", map);
    }
```

