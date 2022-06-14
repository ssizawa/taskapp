$(function(){
    $('#right').on('click', function(){
        $('.sidemenu').animate({
            'display': 'block' 
        }).animate({
            'left': 0
        },
        300
        );
    });
});