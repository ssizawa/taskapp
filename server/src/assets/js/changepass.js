//sidemenu
$(function(){
  let duration = 300;
  let $trriger = $('#trigger');
  let $sidemenu = $('#sidemenu');
  let $trriger_close = $('#trigger-close');

  $trriger.on('click', function(){
    if(!$trriger_close.hasClass('open')){
        $sidemenu.stop(true).animate({
          left: 0
        }, duration);
        $trriger_close.addClass('open');
    }
  });

  $trriger_close.on('click', function(){
    $sidemenu.stop(true).animate({
      left: '-300px'
    }, duration);
    $(this).removeClass('open');
  });
});
  
//Home
$(function(){
  $('#side-item1').on('click', function(){
    location.href='/taskapp';
  });
});

//Settings
$(function(){
  $('#side-item2').on('click', function(){
    location.href='taskapp/settings';
  });
});