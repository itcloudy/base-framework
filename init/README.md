# 初始化数据


## api数据
```yml
apis:
  # 首页
  - name: 管理员首页
    address: /auth/dashboard/num_admin
    method: GET
  - name: 商家首页
    address: /auth/dashboard/num_shop
    method: GET
  - name: 用户信息获取
    address: /auth/user/:id
    method: GET
```


## 菜单
```yml
menus:
  - name: 首页
    unique_tag: admin-home
    route: /dashboard/home
    component: Home
    icon: home-icon
    sequence: 1
  - name: 管理
    unique_tag: companyManage
    route:
    component:
    icon: company-icon
    sequence: 3
    children:
      - name:  审核
        unique_tag: companyAudit
        route: /dashboard/companyAudit
        icon: company-audit-icon
        sequence: 1
        component: CompanyAudit
```