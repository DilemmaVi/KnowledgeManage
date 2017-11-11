$(function () {
   $(document).ready(function () { 

       var E = window.wangEditor
       var editor = new E('#editor')
       // 或者 var editor = new E( document.getElementById('#editor') )
       editor.customConfig.zIndex = 1
       editor.create()

      $('#input-data').click(function (e) {
          e.preventDefault(); //加上这句就可以实现当前选中的li高亮显示，但同时也阻止了href中的页面跳转行为
          var Yjfl=$("#yjfl").val();
          var Ejfl=$("#ejfl").val();
          var Sjfl=$("#sjfl").val();
          var Title=$("#Title").val();
          var Keyword=$("#Keyword").val();
          var Content=editor.txt.text();
          var Contenthtml=editor.txt.html();

          $.ajax({

              type: "post",

              async: true, //同步执行

              url: "/CreateKnowledge",

              data: { "Id":5,"Yjfl":Yjfl,"Ejfl":Ejfl,"Sjfl":Sjfl,"Title":Title,"Content":Content,"Contenthtml":Contenthtml,"Keyword":Keyword},

              dataType: "json", //返回数据形式为json

              success: function (result) {
               if(result.code==200){
                    swal(
                          '恭喜',
                          '提交成功',
                          'success'
                        )
               }else{
                     swal(
                          '出错了',
                          result.reason,
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
   });  
})

