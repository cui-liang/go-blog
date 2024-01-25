import $ from 'jquery';
import 'bootstrap/dist/css/bootstrap.min.css';
import 'bootstrap/dist/js/bootstrap.bundle.min.js';
import '@fortawesome/fontawesome-free/css/all.min.css';

// 引入 highlight.js
import hljs from 'highlight.js';

// 引入本地文件
import '../css/site.scss';

$(document).ready(function () {
    const themeSwitch = $('#theme-switch');
    const html = $('html');
    const body = $('body');
    const main = $('main');

    // 监听主题切换
    themeSwitch.on('change', function () {
        if (themeSwitch.prop('checked')) {
            setTheme('dark');
        } else {
            setTheme('light');
        }
    });

    function setTheme(theme) {
        const isDarkTheme = theme === 'dark';

        html.attr('data-bs-theme', theme);
        const expires = new Date();
        expires.setTime(expires.getTime() + (365 * 24 * 60 * 60 * 1000));
        document.cookie = `theme=${theme};expires=${expires.toUTCString()};path=/`;

        const bodyClasses = isDarkTheme ? ['bg-light', 'bg-dark'] : ['bg-dark', 'bg-light'];
        const mainClasses = isDarkTheme ? ['bg-white', 'bg-black'] : ['bg-black', 'bg-white'];

        body.removeClass(...bodyClasses).addClass(...bodyClasses.reverse());
        main.removeClass(...mainClasses).addClass(...mainClasses.reverse());

        $('#hljs-theme').attr('href', `/dist/${theme}.css`);
    }

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
    body.append('<div id="blueimp-gallery" class="blueimp-gallery blueimp-gallery-controls"><div class="slides"></div><h3 class="title"></h3><a class="prev">‹</a><a class="next">›</a><a class="close">×</a><a class="play-pause"></a><ol class="indicator"></ol></div>');
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