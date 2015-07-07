$(function(){ 

    $(window).on('scroll',function(){
        if($(window).scrollTop()>500){
            $(".top_btn").css("display","block");
        }
        else{
            $(".top_btn").css("display","none");
        }
    }) 

     $("#border").bind("click",function(){
        $(".wraptop").slideToggle();
     });  


    //限制字符个数
     var Sys = {};
        var ua = navigator.userAgent.toLowerCase();
        var s;
        (s = ua.match(/msie ([\d.]+)/)) ? Sys.ie = s[1] :
        (s = ua.match(/firefox\/([\d.]+)/)) ? Sys.firefox = s[1] :
        (s = ua.match(/chrome\/([\d.]+)/)) ? Sys.chrome = s[1] :
        (s = ua.match(/opera.([\d.]+)/)) ? Sys.opera = s[1] :
        (s = ua.match(/version\/([\d.]+).*safari/)) ? Sys.safari = s[1] : 0;



       if(Sys.ie==undefined) {
      
      $(".beyond").each(function(){
        var maxwidth=20;
        if($(this).text().length>maxwidth){
            $(this).text($(this).text().substring(0,maxwidth));
            $(this).html($(this).html()+'...');
        }
    });
  }else{/*兼容IE*/
  
   var wordLimit=function(){
    $(".beyond").each(function(){
        var copyThis = $(this.cloneNode(true)).hide().css({
            'position': 'absolute',
            'width': 'auto',
            'overflow': 'visible'
        }); 
        $(this).after(copyThis);
        if(copyThis.width()>$(this).width()){
            $(this).text($(this).text().substring(0,$(this).html().length-4));
            $(this).html($(this).html()+'...');
            copyThis.remove();
            wordLimit();
        }else{
            copyThis.remove(); //清除复制
            return;
        }
    });
}
wordLimit();
}
  
});

/*轮播控制*/
/*$(document).ready(function(e) {
    var unslider04 = $('#b04').unslider({
        dots: true
    }),
    data04 = unslider04.data('unslider');
     
    $('.unslider-arrow04').click(function() {
        var fn = this.className.split(' ')[1];
        data04[fn]();
    });
});


  */