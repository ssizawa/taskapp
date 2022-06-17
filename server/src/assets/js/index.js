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

//draggable
$(function(){
  let task = '[id^=task-]';
  let box = '[id^=box]';

  $(task).draggable({
    snap: box,
    snapMode: 'inner',
    revert: 'invalid',
    opacity: 0.8,
    cursor: 'move'
  });

  $(box).droppable({
    classes: {
      "ui-droppable-active": "ui-state-active",
      "ui-droppable-hover": "ui-state-hover"
    },
    drop: function(event, ui) {
     var $obj = $(ui.draggable[0]);
     $obj.offset($(this).offset());
     return false;
    }
  });

  $(function(){
    $('.sortable').sortable();
    $(task).draggable({
      connectToSortable: '.sortable'
    })
  });
});