<html lang="zh-CN">

<head>

    <title>{{.title}}-dreamer知识库</title>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="icon" href="http://www.dreamer.cn/favicon.ico" type="image/x-icon">
    <link rel="shortcut icon" href="favicon.ico" mce_href="favicon.ico" type="image/x-icon">
    <script type="text/javascript" src="/static/js/jquery-1.8.0.min.js"></script>
    <link href="/static/css/bootstrap/css/bootstrap.min.css" rel="stylesheet">
    <link href="/static/css/bootstrap/css/bootstrap-theme.min.css" rel="stylesheet">
    <script src="/static/css/bootstrap/js/bootstrap.min.js"></script>
    <script src="/static/css/bootstrap/respond.min.js"></script>
    <link href="/static/css/web-base.css" rel="stylesheet">
    <link href="/static/css/web-black.css" rel="stylesheet">
    <link href="/static/css/kindeditor/editInner.css" rel="stylesheet">
    <script type="text/javascript" src="/static/js/wcpTypes.js"></script>
    <link href="/static/css/nice-select.css" rel="stylesheet">
    <link href="/static/css/sweetalert2.min.css" rel="stylesheet">
    <script type="text/javascript" src="/static/js/sweetalert2.min.js"></script>

    <link href="/static/css/bootstrap-table/bootstrap-table.css" rel="stylesheet">
    <script type="text/javascript" src="/static/css/bootstrap-table/bootstrap-table.js"></script>
    <script type="text/javascript" src="/static/css/bootstrap-table/bootstrap-table-zh-CN.min.js"></script>





</head>

<body>
    <!-- class="navbar navbar-default|navbar-inverse" -->
    <div class="navbar navbar-inverse navbar-fixed-top " role="navigation" style="margin: 0px;">
        <div class="container">
            <!-- Brand and toggle get grouped for better mobile display -->
            <div class="navbar-header">
                </button>
                <a class="navbar-brand" style="color: #ffffff; font-weight: bold; padding: 5px;" href="/search">
                    <img src="http://www.dreamer.cn/favicon.ico" height="40" alt="WCP" title="WCP" align="middle">
                </a>
            </div>
            <!-- Collect the nav links, forms, and other content for toggling -->
            <div class="collapse navbar-collapse" id="bs-example-navbar-collapse-1">
                <ul class="nav navbar-nav">

                    <li>
                        <a href="/list">
                            <span class="glyphicon glyphicon-home"></span> 知识列表</a>
                    </li>


                    <li>
                        <a href="/search">
                            <span class="glyphicon glyphicon-question-sign"></span> 检索知识</a>
                    </li>

                    <li>
                        <a href="/CreateKnowledge">
                            <span class="glyphicon glyphicon-book"></span>创建知识</a>
                    </li>


                    <li class="hidden-xs hidden-sm hidden-md">
                        <a href="/ClassifyManage">
                            <span class="glyphicon glyphicon-tree-conifer"></span> 分类管理</a>
                    </li>


                    <li class="hidden-xs hidden-sm hidden-md">
                        <a href="/BackupData">
                            <span class="glyphicon glyphicon-list-alt"></span> 备份知识 </a>
                    </li>
                                        <li class="hidden-xs hidden-sm hidden-md">
                        <a href="/manage">
                            <span class="glyphicon glyphicon-cog"></span> 管理后台 </a>
                    </li>
                </ul>

   <ul class="nav navbar-nav navbar-right" style="margin-right: 10px;">

    <!-- 登录注销 -->
        {{if eq .Member ""}}
        <li><a href="/login"><span class="glyphicon glyphicon glyphicon-user"></span> 登录</a></li>
        {{else}}
        <li><a href="/login"><span class="glyphicon glyphicon glyphicon-user"></span> {{.Member}}</a></li>
        {{end}}
    
    
    
</ul>
            </div>
        </div>
        <!-- /.navbar-collapse -->
    </div>


    <div class="containerbox">
        <div class="container">
         <div style="margin-top: 30px;"></div>
            {{.LayoutContent}}
        </div>
    </div>
    <div style="margin-top: 10px;"></div>

    <div class="foot">
        Copyright © Dreamer 京ICP备15042091号-2 server@dreamer.cn
    </div>
    <script type="text/javascript">
        $(function () {
            $(window).resize(function () {
                $('.containerbox').css('min-height', $(window).height() - 50);
            });
            $('.containerbox').css('min-height', $(window).height() - 50);
        });
    </script>


</body>

</html>