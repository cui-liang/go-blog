import $ from 'jquery';
import 'bootstrap/dist/css/bootstrap.min.css';
import 'bootstrap/dist/js/bootstrap.bundle.min.js';
import '@fortawesome/fontawesome-free/css/all.min.css';
import 'blueimp-gallery/js/jquery.blueimp-gallery.min';
import 'blueimp-gallery/css/blueimp-gallery.min.css';

// 引入 highlight.js
import hljs from 'highlight.js';

// 引入本地文件
import '../css/main.scss';

$(document).ready(function () {
    // 查找所有代码块并应用语法高亮
    $('pre code').each(function(i, block) {
        hljs.highlightElement(block);
    });

    // back-to-top
    $(window).scroll(function () {
        if ($(this).scrollTop() > 500) {
            $('.back-to-top').fadeIn();
        } else {
            $('.back-to-top').fadeOut();
        }
    });

    $('.back-to-top').click(function (e) {
        e.preventDefault();
        $('html,body').animate({scrollTop: 0}, 'slow');
    });

    // reply
    $(document).on('click', '.reply', function () {
        $('#comment').appendTo($(this).parents('li > .media-body'));
        $('#comment-parent_id').val($(this).parents('li').attr('data-key'));
        $('.md-editor').remove();
        if ($(this).parents('div.media').length > 0) {
            $('#comment-content').val('@' + $(this).closest('.media-body').find('a').first().html() + ' ').mdEditor();
        } else {
            $('#comment-content').val('').mdEditor();
        }
    });

    //gallery
    $('p img,.card-img').each(function(){
        var img_src = $(this).attr('src');
        var img_alt = $(this).attr('alt');
        var reg = /^\/uploads\/images\/\w/;
        if(reg.test(img_src)) {
            img_src = img_src.replace('_thumb.', '.');
        }
        $(this).wrap('<a href="' + img_src + '" title="' + img_alt + '" data-gallery></a>');
    });
});