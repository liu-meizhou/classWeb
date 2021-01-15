## 课程管理

-- doc `需求`

-- goweb `后端代码`

-- vue-element-admin `前端代码`

### 功能

#### 学生

- [x] 查看选课列表(分页还需改进)
- [x] 选课(专业选修)
- [x] 查看个人课表
- [x] 查看课程的分数，绩点

#### 任课老师、系主任

- [x] 查看个人课表
- [x] 查看课程学生和其成绩列表(分页还需改进)
- [x] 修改修改自己课程的学生的成绩(批量登记成绩)
- [ ] 申请功能(给系主任审批)
  - [ ] 申请开课(添加/修改课程)
  - [ ] 申请开设课组(开设者默认为负责老师)

#### 系主任

- [x] 查看所有学生、班级、课程、老师的详细信息(and)
  - [x] CRU基础信息(R分为ReadOne和ReadList，ReadList要带上请求参数(还需改进)，如pageInfo)
    - [x] 学生
    - [x] 班级
    - [x] 课程
    - [x] 老师
  - [x] 查看学生、老师课表
- [x] 班级
  - [x] 班级课表查询
- [ ] 学生
  - [ ] 根据excel表批量导入学生
  - [ ] 排序查询
    - [x] 班级内学生绩点排名查询
    - [ ] 院级内学生绩点排名查询
    - [ ] 校级内学生绩点排序查询
- [x] 课程、班级
  - [x] 查看课程的班级
  - [x] 给班级统一选课
- [ ] 课程、老师
  - [ ] 审批老师开课
  - [ ] 审批老师开设的课组

#### 其他

- [ ] 自动排课(or)
  - [ ] 排课算法
  - [ ] 手动排课
  - [ ] 调课
- [x] 测试接口

### 部署

#### 命令部署

- 条件
  - 电脑需要安装git、docker、docker-compose
- 步骤
  1. `git clone https://github.com/liu-meizhou/classWeb.git`
  2. `cd classWeb`
  3. 进行配置文件修改 `.env`文件 和 `goweb/conf/app.conf`,一般默认也就可以了
  4. `docker-compose up -d`
  5. 等待一段时间，可以用`docker logs xxx`查看是否启动完成
  6. 当后端`go_web`启动完成 在本机浏览器访问 `http://localhost:8080/admin/v1/genDB?token=MrLiu`即可初始化数据库
  7. 当前端`vue_admin_nginx`启动完成 在本机浏览器访问 `http://localhost:9999` 即可到主界面
- 命令更新:
  - 更新前端：`docker-compose up -d vue_admin_nginx` 或者 `docker-compose restart vue_admin_env vue_admin_nginx`
  - 更新后端：`docker-compose restart go_web`  或者  `docker restart go_web`
  - 更新数据库：`docker-compose restart postgresql` 或者  `docker restart postgresql`

#### 自动化部署

- Jenkins
- Github Action
- Drone(需要有一个外网地址，如没有可以使用Gogs自建参考)

##### Jenkins自动化部署

- 条件
  - 电脑需要安装git、docker、docker-compose
  - docker容器需要启动Jenkins,这里提供一套Jenkins的启动方案
    1. `git clone https://github.com/liu-meizhou/classWeb.git`(建议clone到`~/project/`)
    2. `cd classWeb`
    3. `docker-compose -f docker-compose-jenkins.yml up -d`
  - 需要当前目录在`~/project`下
- 步骤
  1. `git clone https://github.com/liu-meizhou/classWeb.git`
  2. 在浏览器打开[Jenkins](http://localhost:8888/)页面 -> 登录 -> 打开[blue ocean](http://localhost:8888/blue/organizations/jenkins/pipelines) -> 选择git[创建pipeline](http://localhost:8888/blue/organizations/jenkins/create-pipeline) ->

##### Github Action





