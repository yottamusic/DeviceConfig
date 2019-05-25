function finish_loading(){
  $('.body-base').css('opacity', '');

  setTimeout(function(){
    $('.loading-screen').animate({
      opacity: 0
    }, 500, function(){
      $(this).remove();
      $('.no-overflow').removeClass('no-overflow');
    })
  }, 2000);
}


function events() {
  $(".drawer").on("click", function () {
    var nav = $(".app-navigation");
    if (nav.hasClass('open')) {
      nav.removeClass('open');
      $('.close-nav, .menu-item').removeClass('visible');
    } else {
      nav.addClass('open');
      $('.close-nav, .menu-item').addClass('visible');
    }
  });

  $(".close-nav").on("click", function () {
    var nav = $(".app-navigation");
    if (nav.hasClass('open')) {
      nav.removeClass('open');
      $('.close-nav, .menu-item').removeClass('visible');
    } else {
      nav.addClass('open');
      $('.close-nav, .menu-item').addClass('visible');
    }
  });

  $("[data-href]").on('click', function(){
    var current_url = window.location.href;
    current_url = current_url.replace(/(\w|\-)*\.html/gi, "");
    var new_url = current_url + $(this).data('href') +  '.html';

    barba.go(new_url);
  });
}


// init Barba with a default "opacity" transition
barba.init({
  transitions: [{
    leave: function (data) {
      var done = this.async();
      $(data.current.container).animate({
        opacity: 0
      }, 250, function(){
        done();
      });
    },
    enter: function (data) {
      var done = this.async();
      $(data.next.container).animate({
        opacity: 1
      }, 250, function(){
        done();
        events();
      });
    }
  }]
});


if (window.addEventListener) {
  window.addEventListener('load', function () {
    try{ finish_loading() }catch(err){}
    try{ events() }catch(err){}
  });
} else {
  window.attachEvent('onload', function () {
    try{ finish_loading() }catch(err){}
    try{ events() }catch(err){}
  });
}
