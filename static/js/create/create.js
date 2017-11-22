$(function () {
    $(document).ready(function () {

        var E = window.wangEditor
        var editor = new E('#editor')
        // 或者 var editor = new E( document.getElementById('#editor') )
        editor.customConfig.zIndex = 1
        editor.create()

        $('#input-data').click(function (e) {
            e.preventDefault(); //加上这句就可以实现当前选中的li高亮显示，但同时也阻止了href中的页面跳转行为
            var Yjfl = $("#yjfl").val();
            var Ejfl = $("#ejfl").val();
            var Sjfl = $("#sjfl").val();
            var Title = $("#Title").val();
            var Keyword = $("#Keyword").val();
            var Content = editor.txt.text();
            var Contenthtml = editor.txt.html();
            if (yjfl=="一级分类" || Ejfl=="二级分类"||Sjfl=="三级分类"){
                swal(
                    '出错了',
                    "请选择完整的分类",
                    'error'
                )   
                return            
            }

            if(Title==""){
                swal(
                    '出错了',
                    "标题不能为空",
                    'error'
                )   
                return                   
            }
            if(Content==""){
                swal(
                    '出错了',
                    "正文不能为空",
                    'error'
                )   
                return                   
            }
            
            
            $.ajax({

                type: "post",

                async: true, //同步执行

                url: "/CreateKnowledge/add",

                data: { "yjfl": Yjfl, "ejfl": Ejfl, "sjfl": Sjfl, "title": Title, "content": Content, "contenthtml": Contenthtml, "keyword": Keyword },

                dataType: "json", //返回数据形式为json

                success: function (result) {
                    if (result.msg == "ok") {
                        swal(
                            '恭喜',
                            '提交成功',
                            'success'
                        )
                    } else {
                        swal(
                            '出错了',
                            result.msg,
                            'error'
                        )
                    }
                },
                error: function (errorMsg) {
                    swal(
                        '出错',
                        errorMsg,
                        'error'
                    )
                }

            });

        });


        $("#yjfl").change(function () {
            var yjfl = $("#yjfl").val();
            $.ajax({

                type: "post",

                url: "/CreateKnowledge/find",

                async: true, //同步执行

                data: { "rank":0,"classify":yjfl },

                dataType: "json", //返回数据形式为json

                success: function (result) {
                    if (result.msg == "ok") {
                        $("#ejfl").empty();
                        $("#sjfl").empty();
                        $("#sjfl").append("<option>三级分类</option>")
                        $.each(result.data,function(index,content){
                            $("#ejfl").append("<option>"+content+"</option>")
                        });
                        $("#ejfl").niceSelect('update'); 
                        $("#sjfl").niceSelect('update');   
                    } else {
                        swal(
                            '出错了',
                            result.msg,
                            'error'
                        )
                    }
                },
                error: function (errorMsg) {
                    swal(
                        '出错',
                        errorMsg,
                        'error'
                    )
                }

            });
        });

        $("#ejfl").change(function () {
            var ejfl = $("#ejfl").val();
            $.ajax({

                type: "post",

                url: "/CreateKnowledge/find",

                async: true, //同步执行

                data: { "rank":2,"classify":ejfl },

                dataType: "json", //返回数据形式为json

                success: function (result) {
                    if (result.msg == "ok") {
                        $("#sjfl").empty();
                        $.each(result.data,function(index,content){
                            $("#sjfl").append("<option>"+content+"</option>")
                        });
                        $("#sjfl").niceSelect('update');   
                    } else {
                        swal(
                            '出错了',
                            result.msg,
                            'error'
                        )
                    }
                },
                error: function (errorMsg) {
                    swal(
                        '出错',
                        errorMsg,
                        'error'
                    )
                }

            });
        });

        $("#preview-data").click(function(){
            swal(
                '温馨提示',
                '功能开发中',
                'warning'
            )
        })


    });
})

