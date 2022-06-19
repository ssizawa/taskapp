//new button
$(function(){
  $('#new-button').on('mouseover', function(){
    $(this).css({
      'background-color': '#e7a41c' 
    });
  });

  $('#new-button').on('mouseout', function(){
    $(this).css({
      'background-color': '#EDAA28' 
    });
  });
});

//trigger
$(function(){
  $('#trigger').on('mouseover', function(){
    $('#right').css({
      'color': '#d19c32' 
    });
  });

  $('#trigger').on('mouseout', function(){
    $('#right').css({
      'color': '#EDAA28' 
    });
  });
});

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

//close
$(function(){
  $('#trigger-close').on('mouseover', function(){
    $('#left').css({
      'color': '#d19c32' 
    });
  });

  $('#trigger-close').on('mouseout', function(){
    $('#left').css({
      'color': '#EDAA28' 
    });
  });
});

//Home
$(function(){
  $('#side-item1').on('click', function(){
    location.reload();
  });
});

//Settings
$(function(){
  $('#side-item2').on('click', function(){
    location.href='taskapp/settings';
  });
});

//draggable
$(function(){
  $('.droppable').sortable({
    connectWith: '.droppable',
    items: '[id^=task-]'
  });
});

//modal
$(function(){
  $('#taskModal').on('hidden.bs.modal', function(){
    $(this).find('form').trigger('reset');
  })
});