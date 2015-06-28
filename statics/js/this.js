$(function(){
   $(".pt").bind("mouseover",function(event){
   		$(".pt").css("opacity","0.5");
         $(".pt span").show();
         event.stopPropagation();
   });
   $(".pt").bind("mouseout",function(event){
        	$(".pt").css("opacity","1");
         $(".pt span").hide();
         event.stopPropagation();/*冒泡有问题*/
   })

})


$(document).ready(function(){
    //限制字符个数
    $(".beyond").each(function(){
        var maxwidth=50;
        if($(this).text().length>maxwidth){
            $(this).text($(this).text().substring(0,maxwidth));
            $(this).html($(this).html()+'...');
        }
    });
});