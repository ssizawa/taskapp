//sidemenu
$(function(){
  var duration = 300;
  var $trriger = $('#trigger');
  var $sidemenu = $('#sidemenu');
  var $trriger_close = $('#trigger-close');

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