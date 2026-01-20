// ==UserScript==
// @name         通用网页允许复制/右键
// @namespace    http://tampermonkey.net/
// @version      1.0
// @description  解除网页禁止复制、剪切、选择文本、右键菜单的限制
// @author       GitHub Copilot
// @match        *://*/*
// @grant        GM_addStyle
// @run-at       document-start
// ==/UserScript==

(function() {
    'use strict';

    // 1. 注入 CSS 强制允许用户选择文本
    const css = `
        * {
            -webkit-user-select: text !important;
            -moz-user-select: text !important;
            -ms-user-select: text !important;
            user-select: text !important;
        }
    `;

    // 尝试注入样式
    try {
        if (typeof GM_addStyle !== 'undefined') {
            GM_addStyle(css);
        } else {
            const style = document.createElement('style');
            style.innerHTML = css;
            (document.head || document.documentElement).appendChild(style);
        }
    } catch (err) {
        console.log('注入样式失败:', err);
    }

    // 2. 需要解除限制的事件列表
    const events = ['copy', 'cut', 'contextmenu', 'selectstart', 'dragstart', 'mousedown', 'mouseup'];

    // 3. 在捕获阶段拦截事件
    events.forEach(evt => {
        window.addEventListener(evt, function(e) {
            // 注意：这是一种较"暴力"的解锁方式。
            // 它的原理是：事件在捕获阶段就停止传播，这样网页上定义的"禁止复制"脚本（通常在冒泡阶段）就不会被触发。
            // 副作用：可能会让一些复杂的 Web 应用（如在线文档编辑器、右键有自定义菜单的网站）功能失效。
            // 如果遇到正常交互受阻，请在插件菜单中临时关闭此脚本。
            e.stopImmediatePropagation();
        }, true);
    });

    // 4. 清除特定 DOM 属性上的限制（例如 <body oncopy="return false">）
    function clearDomEvents() {
        const doc = document;
        const body = document.body;
        
        events.forEach(evt => {
            const onEvt = 'on' + evt;
            // 清除 document 上的属性
            if (doc && doc[onEvt]) doc[onEvt] = null;
            // 清除 body 上的属性
            if (body && body[onEvt]) body[onEvt] = null;
        });
    }

    // 初始化时清除
    clearDomEvents();
    
    // 定时检查并清除，防止页面动态加载脚本后重新加上限制
    setInterval(clearDomEvents, 2000);

})();
