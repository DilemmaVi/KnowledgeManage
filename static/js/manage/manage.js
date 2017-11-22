$(function () {
    $(document).ready(function () {
        $("#yjfl").niceSelect();
        $("#ejfl").niceSelect();
        $("#sjfl").niceSelect();

        $("#btn_add").click(function () {

            swal({
                title: '注册账户',
                type: 'info',
                html:
                    '<form id="register_form">\n' +
                    '<div class="input-group">\n' +
                    '    <span class="input-group-addon" >用户名</span>\n' +
                    '    <input type="text" id="add_account" name="account" class="form-control" placeholder="用户名" aria-describedby="basic-addon1">\n' +
                    '</div>\n' +
                    '<div style="padding: 5px 0; color: #ccc"></div>\n' +
                    '<div class="input-group">\n' +
                    '    <span class="input-group-addon" >密码</span>\n' +
                    '    <input type="password" id="add_password" name="password" class="form-control" placeholder="密码" aria-describedby="basic-addon1">\n' +
                    '</div>\n' +
                    '<div style="padding: 5px 0; color: #ccc"></div>\n' +
                    '<div class="input-group">\n' +
                    '    <span class="input-group-addon" >确认密码</span>\n' +
                    '    <input type="password" id="password_confirm"  class="form-control" placeholder="确认密码" aria-describedby="basic-addon1">\n' +
                    '</div>\n' +
                    '<div style="padding: 5px 0; color: #ccc"></div>\n' +
                    '<div class="input-group">\n' +
                    '    <span class="input-group-addon" >姓名</span>\n' +
                    '    <input type="text" id="add_name" name="name" class="form-control" placeholder="姓名" aria-describedby="basic-addon1">\n' +
                    '</div>\n' +
                    '<div style="padding: 5px 0; color: #ccc"></div>\n' +
                    '<div class="input-group">\n' +
                    '    <span class="input-group-addon" >手机号</span>\n' +
                    '    <input type="text" id="add_phone" name="phone" class="form-control" placeholder="手机号" aria-describedby="basic-addon1">\n' +
                    '</div>\n' +
                    '<div style="padding: 5px 0; color: #ccc"></div>\n' +
                    '<div class="input-group">\n' +
                    '    <span class="input-group-addon" >邮箱</span>\n' +
                    '    <input type="text" id="add_email" name="email" class="form-control" placeholder="邮箱" aria-describedby="basic-addon1">\n' +
                    '</div>\n' +
                    '<div style="padding: 5px 0; color: #ccc"></div>\n' +
                    '<div class="input-group">\n' +
                    '    <span class="input-group-addon">角色</span>\n' +
                    '    <select id="add_role" name="role" class="wide small">\n' +
                    '    <option value="2">普通用户</option>\n' +
                    '    <option value="1">管理员</option>\n' +
                    '<option value="0" >超级管理员</option>\n' +
                    '        </select>\n' +
                    '</div>\n' +
                    '<div style="padding: 5px 0; color: #ccc"></div>\n' +
                    '<div class="input-group">\n' +
                    '    <span class="input-group-addon" >是否激活</span>\n' +
                    '    <select id="add_status" name="status" class="wide small">\n' +
                    '    <option value="0">是</option>\n' +
                    '    <option value="1">否</option>\n' +
                    '        </select>\n' +
                    '</div></form>',
                showCancelButton: true,
                allowOutsideClick: false,
                confirmButtonColor: '#3085d6',
                cancelButtonColor: '#d33',
                confirmButtonText: '提交注册',
                cancelButtonText: "取消",
                preConfirm: function () {
                    return new Promise(function (resolve) {
                        var flag = false
                        $("div.input-group>input").each(function () {
                            if ($(this).val() == "") {
                                flag = true
                                reject($(this).attr('placeholder') + "不能为空")
                                return false
                            }
                        })
                        if (flag) {
                            return
                        } else {
                            $.ajax({
                                type: "POST",
                                url: "/manage/add",
                                data: $('#register_form').serialize(),// 你的formid
                                success: function (data) {
                                    if (data.msg != "ok") {
                                        reject(data.msg)
                                    } else {
                                        resolve("0")
                                    }
                                }
                            });
                        }

                    });
                }
            }).then(function (result) {
                if (result) {
                    swal(
                        {
                            title: '恭喜!',
                            text: '大吉大利,添加成功!',
                            type: 'success',
                            timer: 1000
                        }
                    );
                    $('#mytab').bootstrapTable('refresh', { url: '/manage/members' });
                }
            })
            $("#add_role").niceSelect();
            $("#add_status").niceSelect();
        })



        function reject(text) {
            $(".swal2-validationerror").html(text)
            $(".swal2-validationerror").css('display', 'block')

        }


        //根据窗口调整表格高度
        $(window).resize(function () {
            $('#mytab').bootstrapTable('resetView', {
                height: tableHeight()
            })
        })
        //生成用户数据
        $('#mytab').bootstrapTable({
            method: 'post',
            contentType: "application/x-www-form-urlencoded",//必须要有！！！！
            url: "/manage/members",//要请求数据的文件路径
            height: tableHeight(),//高度调整
            toolbar: '#toolbar',//指定工具栏
            striped: true, //是否显示行间隔色
            dataField: "res",//bootstrap table 可以前端分页也可以后端分页，这里
            //我们使用的是后端分页，后端分页时需返回含有total：总记录数,这个键值好像是固定的
            //rows： 记录集合 键值可以修改  dataField 自己定义成自己想要的就好
            pageNumber: 1, //初始化加载第一页，默认第一页
            pagination: true,//是否分页
            queryParamsType: 'limit',//查询参数组织方式
            queryParams: queryParams,//请求服务器时所传的参数
            sidePagination: 'server',//指定服务器端分页
            pageSize: 10,//单页记录数
            pageList: [5, 10, 20, 30],//分页步进值
            showRefresh: false,//刷新按钮
            showColumns: true,
            clickToSelect: true,//是否启用点击选中行
            toolbarAlign: 'right',//工具栏对齐方式
            buttonsAlign: 'right',//按钮对齐方式
            toolbar: '#toolbar',//指定工作栏
            columns: [
                {
                    title: '用户名',
                    field: 'account',
                },
                {
                    title: '姓名',
                    field: 'name',
                },
                {
                    title: '手机号',
                    field: 'phone',
                },
                {
                    title: '邮箱',
                    field: 'email'
                },
                {
                    title: '角色',
                    field: 'role_name',
                },
                {
                    title: '注册日期',
                    field: 'create_time',
                },
                {
                    title: '上次登录日期',
                    field: 'last_login_time',
                },

                {
                    title: '是否激活',
                    field: 'status',
                    align: 'center',
                    //列数据格式化
                    formatter: statusFormatter
                },
                {
                    title: '操作',
                    field: 'id',
                    formatter: operateFormatter
                }
            ],
            locale: 'zh-CN',//中文支持,
            responseHandler: function (res) {
                //在ajax获取到数据，渲染表格之前，修改数据源
                return res;
            }
        })
        //三个参数，value代表该列的值
        function statusFormatter(value, row, index) {
            if (value == "1") {
                return '<i class="glyphicon glyphicon-remove" style="color:red"></i>'
            } else if (value == "0") {
                return '<i class="glyphicon glyphicon-ok" style="color:green"></i>'
            } else {
                return '数据错误'
            }
        }

        function operateFormatter(value, row, index) {
            return '<div class="btn-toolbar" role="toolbar"><div class="btn-group  btn-group-xs" role="group" aria-label="First group"><button type="button" class="btn btn-info" data-id="' + value + '">编辑</button></div><div class="btn-group  btn-group-xs" role="group" aria-label="Second group"><button type="button" class="btn btn-danger" data-id="' + value + '">删除</button></div></div>'
        }

        //请求服务数据时所传参数
        function queryParams(params) {
            return {
                //每页多少条数据
                pageSize: params.limit,
                //请求第几页
                pageIndex: params.pageNumber,
                Name: $('#search_name').val(),
                Tel: $('#search_tel').val(),
                Email: $('#search_email').val(),
            }
        }
        //查询按钮事件
        $('#search_btn').click(function () {
            $('#mytab').bootstrapTable('refresh', { url: '/manage/members' });
        })
        //tableHeight函数
        function tableHeight() {
            //可以根据自己页面情况进行调整
            return $(window).height() - 350;
        }

        //删除按钮事件
        $(document.body).on("click", ".btn-toolbar>div>button.btn.btn-danger", function () {
            var id = $(this).attr('data-id')

            swal({
                title: '确定删除吗?',
                text: "删除后数据将无法恢复!",
                type: 'warning',
                showCancelButton: true,
                confirmButtonColor: '#3085d6',
                cancelButtonColor: '#d33',
                confirmButtonText: '确定删除',
                cancelButtonText: "取消"
            }).then(function (isConfirm) {
                if (isConfirm) {
                    $.ajax({
                        type: "POST",
                        url: "/manage/delete",
                        data: { "id": id },
                        dataType: "json",
                        success: function (data) {
                            if (data.msg == "ok") {
                                swal({
                                    title: '恭喜!',
                                    text: '大吉大利,删除成功!',
                                    type: 'success',
                                    timer: 1000
                                })
                                $('#mytab').bootstrapTable('refresh', { url: '/manage/members' });
                            } else {
                                swal(
                                    '删除失败!',
                                    data.msg,
                                    'error'
                                )
                            }
                        }
                    });

                }
            })

        })

        //编辑按钮事件
        $(document.body).on("click", ".btn-toolbar>div>button.btn.btn-info", function () {
            var id = $(this).attr('data-id')
            $.ajax({
                type: "POST",
                url: "/manage/find",
                data: { "id": id },
                dataType: "json",
                success:function(data){
                    swal({
                        title: '不改密码请留空',
                        type: 'info',
                        html:
                            '<form id="change_form">\n' +
                            '<div class="input-group">\n' +
                            '    <span class="input-group-addon" >密码</span>\n' +
                            '    <input type="password" id="change_password1" name="change_password" class="form-control" placeholder="修改密码" aria-describedby="basic-addon1">\n' +
                            '</div>\n' +
                            '<div style="padding: 5px 0; color: #ccc"></div>\n' +
                            '<div class="input-group">\n' +
                            '    <span class="input-group-addon" >确认密码</span>\n' +
                            '<input type="hidden" name="change_id" value="'+id+'" ><input type="password" id="change_password2" class="form-control" placeholder="确认密码" aria-describedby="basic-addon1">\n' +
                            '</div>\n' +
                            '<div style="padding: 5px 0; color: #ccc"></div>\n' +
                            '<div class="input-group">\n' +
                            '    <span class="input-group-addon" >姓名</span>\n' +
                            '    <input type="text" id="change_name" name="change_name" class="form-control" placeholder="姓名" value="'+data.res.name+'" aria-describedby="basic-addon1">\n' +
                            '</div>\n' +
                            '<div style="padding: 5px 0; color: #ccc"></div>\n' +
                            '<div class="input-group">\n' +
                            '    <span class="input-group-addon" >手机号</span>\n' +
                            '    <input type="text" id="change_phone" name="change_phone" class="form-control" placeholder="手机号"  value="'+data.res.phone+'" aria-describedby="basic-addon1">\n' +
                            '</div>\n' +
                            '<div style="padding: 5px 0; color: #ccc"></div>\n' +
                            '<div class="input-group">\n' +
                            '    <span class="input-group-addon" >邮箱</span>\n' +
                            '    <input type="text" id="change_email" name="change_email" class="form-control" placeholder="邮箱" value="'+data.res.email+'" aria-describedby="basic-addon1">\n' +
                            '</div>\n' +
                            '<div style="padding: 5px 0; color: #ccc"></div>\n' +
                            '<div class="input-group">\n' +
                            '    <span class="input-group-addon">角色</span>\n' +
                            resolveRoleName(data.res.role)+
                            '</div>\n' +
                            '<div style="padding: 5px 0; color: #ccc"></div>\n' +
                            '<div class="input-group">\n' +
                            '    <span class="input-group-addon">是否激活</span>\n' +
                            resolveStatus(data.res.status)+
                            '</div></form>',
                        showCancelButton: true,
                        allowOutsideClick: false,
                        confirmButtonColor: '#3085d6',
                        cancelButtonColor: '#d33',
                        confirmButtonText: '确认修改',
                        cancelButtonText: "取消",
                        preConfirm: function () {
                            return new Promise(function (resolve) {
                                var flag = false
                                $("div.input-group>input").each(function () {
                                    if($("#change_password1").val()!=$("#change_password2").val()){
                                        reject("两次输入的密码不一致")
                                        flag=true
                                        return false
                                    }
                                    if ($(this).val() == "") {
                                        if($(this).attr('placeholder').indexOf("密码") >= 0){
                                            return true
                                        }
                                        flag = true
                                        reject($(this).attr('placeholder') + "不能为空")
                                        return false
                                    }
                                })
                                if (flag) {
                                    return
                                } else {
                                    $.ajax({
                                        type: "POST",
                                        url: "/manage/edit",
                                        data: $('#change_form').serialize(),
                                        success: function (data) {
                                            if (data.msg != "ok") {
                                                reject(data.msg)
                                            } else {
                                                resolve("0")
                                            }
                                        }
                                    });
                                }
        
                            });
                        }
                    }).then(function (result) {
                        if (result) {
                            swal(
                                {
                                    title: '恭喜!',
                                    text: '大吉大利,修改成功!',
                                    type: 'success',
                                    timer: 1000
                                }
                            );
                            $('#mytab').bootstrapTable('refresh', { url: '/manage/members' });
                        }
                    })
                    $("#change_role").niceSelect();
                    $("#change_status").niceSelect();
                }
                
            })

        });

        function resolveRoleName(data){
            if(data==0){
                return '<select id="change_role" name="change_role" class="wide small">\n' +
                '    <option value="2">普通用户</option>\n' +
                '    <option value="1">管理员</option>\n' +
                '<option value="0"  selected = "selected">超级管理员</option>\n' +
                '        </select>\n'
            }else if(data==1){
                return '<select id="change_role" name="change_role" class="wide small">\n' +
                '    <option value="2">普通用户</option>\n' +
                '    <option value="1" selected = "selected" >管理员</option>\n' +
                '<option value="0" >超级管理员</option>\n' +
                '        </select>\n'
            }else{
                return '<select id="change_role" name="change_role" class="wide small">\n' +
                '    <option value="2" selected = "selected">普通用户</option>\n' +
                '    <option value="1"  >管理员</option>\n' +
                '<option value="0" >超级管理员</option>\n' +
                '        </select>\n'               
            }

        }

        function resolveStatus(data){
            if(data==0){
                return '<select id="change_status" name="change_status" class="wide small">\n' +
                '    <option value="0" selected = "selected">是</option>\n' +
                '    <option value="1">否</option>\n' +
                '        </select>\n' 

            }else{
                return '<select id="change_status" name="change_status" class="wide small">\n' +
                '    <option value="0" >是</option>\n' +
                '    <option value="1" selected = "selected">否</option>\n' +
                '        </select>\n'                
            }
        }

    })
})

