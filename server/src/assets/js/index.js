//new button
$(function(){
  $('#new-button').on('mouseover', function(){
    $(this).css({
      'background-color': '#f7d9ab',
    });

    $('#plus').css({
      'color': '#EDAA28',
    });

    $('#new').css({
      'color': '#EDAA28',
    });
  });

  $('#new-button').on('mouseout', function(){
    $(this).css({
      'background-color': '#EDAA28',
    });

    $('#plus').css({
      'color': 'white',
    });

    $('#new').css({
      'color': 'white',
    });
  });
});

//trigger
$(function(){
  $('#trigger').on('mouseover', function(){
    $('#right').css({
      'color': '#f7d9ab' 
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
      'color': '#f7d9ab' 
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
    items: 'div[class^="task-"]'
  });
});

$(function(){
  $('.todo').droppable({
    drop: function(event, ui) {
      var nameTag = ui.draggable.attr('name');
      var dropped_taskName = nameTag.split('&')[0];
      var dropped_taskDescription = nameTag.split('&')[1];

      $.post('update_todo', 'taskName=' + dropped_taskName + '&taskDescription=' + dropped_taskDescription);
      location.reload();
      location.reload();  
    }
  });
});

$(function(){
  $('.doing').droppable({
    drop: function(event, ui) {
      var nameTag = ui.draggable.attr('name');
      var dropped_taskName = nameTag.split('&')[0];
      var dropped_taskDescription = nameTag.split('&')[1];

      $.post('update_doing', 'taskName=' + dropped_taskName + '&taskDescription=' + dropped_taskDescription);
      location.reload();
      location.reload(); 
    }
  }); 
});

$(function(){
  $('.done').droppable({
    drop: function(event, ui) {
      var nameTag = ui.draggable.attr('name');
      var dropped_taskName = nameTag.split('&')[0];
      var dropped_taskDescription = nameTag.split('&')[1];

      $.post('update_done', 'taskName=' + dropped_taskName + '&taskDescription=' + dropped_taskDescription);
      location.reload();
      location.reload(); 
    }
  }); 
});

//task
$(function(){
  $('.task-done').on('click', function(){
    var nameTag = $(this).attr('name');
    var taskName = nameTag.split('&')[0];
    var taskDescription = nameTag.split('&')[1];

    if(confirm('Do you want delete "' + taskName + '"?')){
      $.post('delete_task', 'taskName=' + taskName + '&taskDescription=' + taskDescription);
    }
    location.reload();
    location.reload(); 
  });
});

$(function(){
  $('div[class^="task-"]').on('mouseover', function(){
    $(this).css({
      'background-color': '#ff7171'
    });
  });

  $('div[class^="task-"]').on('mouseout', function(){
    $(this).css({
      'background-color': '#FF4F50'
    });
  });

  $('div[class^="task-"]').on('mousedown', function(){
    $(this).css({
      'border': 'solid 6px orange'
    });
  });

  $('div[class^="task-"]').on('mouseup', function(){
    $(this).css({
      'border-style': 'none'
    });
  });
});

//modal
$(function(){
  $('#taskModal').on('hidden.bs.modal', function(){
    $(this).find('form').trigger('reset');
  });
});