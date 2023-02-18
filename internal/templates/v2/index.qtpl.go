// Code generated by qtc from "index.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

package v2

import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

func StreamIndex(qw422016 *qt422016.Writer, sessionProfile ProfileView, showCharity bool, registrationHistoryProfiles []FullProfileView) {
	qw422016.N().S(`
<!DOCTYPE html>
<html lang="en">
<head>
    <title>GitHub profile views statistic badge</title>
    <meta name="description" content="GitHub profile views statistic badge">
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="author" type="text/plain" href="https://u8views.com/humans.txt"/>

    `)
	streamfavicon(qw422016)
	qw422016.N().S(`

    <link rel="preconnect" href="https://fonts.googleapis.com">
    <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
    <link href="https://fonts.googleapis.com/css2?family=Inter:wght@400;500&display=swap" rel="stylesheet">
    <script src="https://cdn.jsdelivr.net/npm/apexcharts"></script>

    <style>html{box-sizing:border-box}*,:before,:after{box-sizing:inherit}html{-webkit-text-size-adjust:100%;-moz-text-size-adjust:100%;text-size-adjust:100%}body{-webkit-overflow-scrolling:touch}html,body,div,span,applet,object,iframe,h1,h2,h3,h4,h5,h6,p,blockquote,pre,a,abbr,acronym,address,big,cite,code,del,dfn,em,img,ins,kbd,q,s,samp,small,strike,strong,sub,sup,tt,var,b,u,i,center,dl,dt,dd,ol,ul,li,fieldset,form,label,legend,table,caption,tbody,tfoot,thead,tr,th,td,article,aside,canvas,details,embed,figure,figcaption,footer,header,menu,nav,output,ruby,section,summary,time,mark,audio,video{margin:0;padding:0;border:0}article,aside,details,figcaption,figure,footer,header,main,menu,nav,section,summary{display:block}audio,canvas,progress,video{display:inline-block}audio:not([controls]){display:none;height:0}a{background-color:transparent}abbr[title]{border-bottom:none;text-decoration:underline;-webkit-text-decoration:underline dotted;text-decoration:underline dotted}b,strong{font-weight:bolder}dfn{font-style:italic}mark{background-color:#ff0;color:#000}svg:not(:root){overflow:hidden}code,kbd,pre,samp{font-family:monospace,monospace;font-size:1em}hr{box-sizing:content-box;height:0;overflow:visible}button,input,select,textarea{font:inherit;margin:0}button,input{overflow:visible}button,select{text-transform:none}button,[type=button],[type=reset],[type=submit]{-webkit-appearance:button;-moz-appearance:button;appearance:button}input,textarea,button,select,a{-webkit-tap-highlight-color:transparent}address{font-style:normal}a:focus:not(:focus-visible),select:focus:not(:focus-visible),input:focus:not(:focus-visible),textarea:focus:not(:focus-visible){outline:0}button::-moz-focus-inner,[type=button]::-moz-focus-inner,[type=reset]::-moz-focus-inner,[type=submit]::-moz-focus-inner{border-style:none;padding:0}button,input[type=reset],input[type=button],input[type=submit]{cursor:pointer}button[disabled],input[disabled]{cursor:default}button{-webkit-appearance:none;-moz-appearance:none;appearance:none;background:0 0;padding:0;border:0;border-radius:0;line-height:1}button:focus:not(:focus-visible){outline:0}a,a:hover{text-decoration:none}[href="javascript:void();"],[href="javascript:"]{cursor:default}ul,ol{list-style:none}blockquote,q{quotes:none}blockquote:before,blockquote:after,q:before,q:after{content:none}table{border-collapse:collapse;border-spacing:0}input[type=text],input[type=password],input[type=date],input[type=datetime],input[type=datetime-local],input[type=email],input[type=month],input[type=number],input[type=search],input[type=tel],input[type=time],input[type=url],input[type=week],textarea{box-sizing:border-box}[type=checkbox],[type=radio]{box-sizing:border-box;margin:0;padding:0}input[type=search]{-webkit-appearance:textfield;-moz-appearance:textfield;appearance:textfield;outline-offset:-2px}input[type=search]::-webkit-search-decoration,input[type=search]::-webkit-search-cancel-button{-webkit-appearance:none;appearance:none}input[type=number]::-webkit-inner-spin-button,input[type=number]::-webkit-outer-spin-button{height:auto;-webkit-appearance:none;appearance:none}::-webkit-file-upload-button{-webkit-appearance:button;appearance:button;font:inherit}input[type=number]{-webkit-appearance:textfield;-moz-appearance:textfield;appearance:textfield}select{width:100%;height:20px;border:0;background:0 0}textarea{resize:none;border:0;overflow:auto}::-webkit-input-placeholder{color:#777;line-height:normal}::-moz-placeholder{color:#777;line-height:normal}::placeholder{color:#777;line-height:normal}[hidden]{display:none}.headline{font-family:"Inter",arial,helvetica neue,helvetica,sans-serif;font-weight:500;color:#111028;line-height:120%}.headline--lvl1{font-size:44px}.headline--lvl2{font-size:38px}.headline--lvl3{font-size:28px}.headline--lvl4{font-size:18px}.button{font-family:"Inter",arial,helvetica neue,helvetica,sans-serif;font-size:18px;font-weight:500;display:-webkit-inline-flex;display:inline-flex;-webkit-align-items:center;align-items:center;-webkit-justify-content:center;justify-content:center;height:48px;border-radius:4px;padding-left:32px;padding-right:32px;border-width:2px;border-style:solid;border-color:transparent;transition:background-color .225s ease,box-shadow .225s ease,border-color .225s ease,color .225s ease}.button--black{color:#fff;background-color:#111028;border-color:#111028}.button--black:hover{background-color:#3f3a92;border-color:#3f3a92}.button--black:active{background-color:#2f2c6d;border-color:#2f2c6d}.button--black:focus-visible{background-color:#3f3a92;border-color:#dcdbf0;box-shadow:none}.button--black:disabled{color:#79769c;background-color:#dcdbf0;border-color:#dcdbf0}.input{display:-webkit-flex;display:flex;-webkit-flex-direction:column;flex-direction:column}.input.is-active .input__clear{display:-webkit-flex;display:flex}.input.is-error .input__error{display:-webkit-flex;display:flex}.input.is-error .input__element{border-color:#e9331a}.input.is-error .input__error-text{display:block}.input__label{font-size:16px;line-height:19px;color:#8c8ba3}.input__wrapper{position:relative;margin-top:8px}.input__element{font-size:18px;line-height:22px;color:#111028;height:48px;width:100%;padding-left:20px;padding-right:48px;border-radius:4px;border-width:1px;border-style:solid;border-color:transparent;background-color:#f3f3f6;transition:border-color .225s ease,box-shadow .225s ease}.input__element:hover{border-color:#918fc7;box-shadow:0 0 6px rgba(67,66,88,.16)}.input__element:focus{box-shadow:none;outline:0}.input__element::-webkit-input-placeholder{font-size:18px;line-height:22px;color:#adacc2}.input__element::-moz-placeholder{font-size:18px;line-height:22px;color:#adacc2}.input__element::placeholder{font-size:18px;line-height:22px;color:#adacc2}.input__element:disabled{pointer-events:none}.input__element:disabled::-webkit-input-placeholder{color:#d4d4d4}.input__element:disabled::-moz-placeholder{color:#d4d4d4}.input__element:disabled::placeholder{color:#d4d4d4}.input__clear{display:none;position:absolute;right:20px;top:12px;width:24px;height:24px;-webkit-align-items:center;align-items:center;-webkit-justify-content:center;justify-content:center}.input__error{display:none;position:absolute;right:20px;top:12px}.input__error-text{display:none;font-size:14px;line-height:17px;color:#e9331a;position:absolute;top:calc(100% + 8px);left:0}.select__label{font-size:16px;line-height:120%;color:#8c8ba3}.select__wrapper{height:40px;padding-top:8px;padding-bottom:8px;padding-left:15px;padding-right:15px;border-radius:4px;border:1px solid #dcdbf0;margin-top:12px;transition:border-color .225s ease}.select__wrapper:focus-within{border-color:#b6b5d9}.select__element{font-size:18px;line-height:25px;color:#201f3a;cursor:pointer}.select__element option{font-size:16px;color:#7c7b88}.select__element option:disabled{color:#cbcadb}.select__element:focus-visible{outline:0}.checkbox{font-size:18px;color:#3c3a59;display:-webkit-flex;display:flex;-webkit-align-items:center;align-items:center;height:24px;cursor:pointer}.checkbox__input{position:absolute;-webkit-appearance:none;-moz-appearance:none;appearance:none}.checkbox__input:focus-visible{width:18px;height:18px;outline:0;box-shadow:0 0 0 2px #000,0 0 0 3px #fff;border-radius:4px}.checkbox__input:checked+.checkbox__element{background-color:#3c3a59;background-size:17px 14px;background-repeat:no-repeat;background-position:3px center;background-image:url("data:image/svg+xml,%3Csvg width='17' height='14' fill='none' xmlns='http://www.w3.org/2000/svg'%3E%3Cpath fill-rule='evenodd' clip-rule='evenodd' d='M16.173.26a1 1 0 0 1 .067 1.413L5.534 13.449.293 8.207a1 1 0 1 1 1.414-1.414l3.759 3.758L14.76.327A1 1 0 0 1 16.173.26Z' fill='%23fff'/%3E%3C/svg%3E")}.checkbox__element{width:24px;height:24px;border:1px solid #3c3a59;border-radius:4px;background-color:#fff;margin-right:12px;transition:background-color .225s ease}.filters__group{padding-top:32px}.filters__sub-headline{font-size:18px;font-weight:600;color:#3c3a59;height:24px;line-height:24px}.filters__visibility{width:28px;height:28px;position:relative;cursor:pointer}.filters__visibility:focus-visible{outline:0;box-shadow:0 0 0 2px #000,0 0 0 3px #fff;border-radius:4px}.filters__visibility:before{content:"";position:absolute;top:50%;left:50%;width:14px;height:2px;margin-left:-7px;margin-top:-1px;background-color:#828282}.filters__elements{margin-top:16px;max-height:160px;overflow:auto}.filters__elements--no-scroll{overflow:initial}.filters__elements::-webkit-scrollbar{width:4px}.filters__elements::-webkit-scrollbar-track{border-radius:8px;background-color:#dddce4}.filters__elements::-webkit-scrollbar-thumb{border-radius:8px;background-color:#b8b6e2}.filters__filled-elements{display:-webkit-flex;display:flex;-webkit-flex-wrap:wrap;flex-wrap:wrap;gap:8px;margin-top:16px}.filters__element:not(:first-child){margin-top:16px}html{height:100%}body{font-family:"Inter",arial,helvetica neue,helvetica,sans-serif;font-size:16px;color:#21212b;display:grid;grid-template-rows:auto 1fr auto;min-height:100%;background-color:#f2f2f3}.wrapper{max-width:1248px;width:100%;padding-left:16px;padding-right:16px;margin-left:auto;margin-right:auto}.disabled-scroll{overflow-y:hidden}.visually-hidden:not(:focus):not(:active){position:absolute;width:1px;height:1px;margin:-1px;border:0;padding:0;white-space:nowrap;-webkit-clip-path:inset(100%);clip-path:inset(100%);clip:rect(0 0 0 0);overflow:hidden}.header__wrapper{display:-webkit-flex;display:flex;-webkit-align-items:center;align-items:center}.header__search{width:100%;padding:40px;border-radius:24px;background-color:#fff;margin-top:80px}.header__action-group{display:-webkit-flex;display:flex;-webkit-align-items:flex-end;align-items:flex-end;-webkit-justify-content:space-between;justify-content:space-between;margin-top:32px}.header__input{margin-top:32px}.footer__group{color:rgba(0,0,0,.8);display:grid;grid-template-columns:1fr auto auto;-webkit-align-items:start;align-items:start;padding-top:40px;padding-bottom:60px;padding-left:0;padding-right:0;border-top:1px solid #d0cfdf}.footer__copyrights{font-size:14px;line-height:20px;color:#585764;display:-webkit-flex;display:flex;-webkit-align-items:center;align-items:center;-webkit-column-gap:14px;-moz-column-gap:14px;column-gap:14px}.footer__dev{display:-webkit-flex;display:flex;-webkit-flex-direction:column;flex-direction:column;margin-right:95px}.footer__dev:last-child{margin-right:45px}.footer__dev-label{font-size:12px;line-height:17px;color:#9f9ead}.footer__dev-name{font-size:14px;line-height:20px;color:#585764;margin-top:8px}.footer__dev-link{font-size:14px;line-height:20px;color:#3c2df9;-webkit-align-self:flex-start;align-self:flex-start;margin-top:4px}</style>
    <style>.footer{background:#020111}.footer__wrapper{padding:104px 0 24px}.footer__group-u8{display:-webkit-flex;display:flex;-webkit-justify-content:space-between;justify-content:space-between}.footer__info{color:#fff}.footer__title{display:block;font-weight:600;font-size:24px;line-height:120%;margin-bottom:16px;color:#fff}.footer__subtitle{line-height:140%;margin-bottom:55px;width:275px;color:#c6cad2}.footer__map{margin-bottom:56px}.footer__copyrights{display:-webkit-flex;display:flex;gap:14px;color:#cfd2d7;font-size:14px;border-top:1px solid #262a30;padding-top:16px}.footer__middle-section{display:-webkit-flex;display:flex;-webkit-flex-direction:column;flex-direction:column;color:#fff;font-size:18px;width:300px}.footer__statistics-link{margin-bottom:24px;color:#fff}.footer__statistics-link:hover{font-weight:600}.footer__github-link{color:#fff;margin-bottom:32px;display:-webkit-flex;display:flex;-webkit-align-items:center;align-items:center;gap:15px}.footer__github-link:hover{font-weight:600}.footer__support{display:-webkit-flex;display:flex;-webkit-flex-direction:column;flex-direction:column;gap:16px;background:0 0}.footer__link{padding:16px 32px 16px 40px;border:1px solid #636077;border-radius:4px;color:#fff;transition:border .2s;width:100%}.footer__link:hover{border:1px solid #fff}.footer__link:active{border:1px solid #fff;background:#060423}.footer__figure{display:-webkit-flex;display:flex;gap:16px}.header{box-shadow:0 4px 12px rgba(134,132,177,.1);background:#fff;display:-webkit-flex;display:flex;-webkit-align-items:center;align-items:center;height:64px}.header__wrapper{display:-webkit-flex;display:flex;-webkit-justify-content:space-between;justify-content:space-between;-webkit-align-items:center;align-items:center;width:1216px;height:100%}.header__logo{line-height:0;display:block;margin-right:auto}.header__profile{position:relative}.header__stars{margin-right:32px;display:-webkit-flex;display:flex;-webkit-align-items:center;align-items:center}.modal.active{opacity:1;pointer-events:all}.modal{position:absolute;display:-webkit-flex;display:flex;-webkit-flex-direction:column;flex-direction:column;width:540px;right:-57%;top:166%;border-radius:2px;background:#fff;opacity:0;pointer-events:none}.modal::after{content:"";position:absolute;right:40px;top:-11px;width:0;height:0;border-left:12px solid transparent;border-right:12px solid transparent;border-bottom:12px solid #fff;border-radius:1px}.modal__profile{display:-webkit-flex;display:flex;gap:16px;padding:24px;border-bottom:1px solid #dddce4}.modal__profile-info{width:100%;padding-top:15px}.modal__profile-name{font-weight:600;font-size:22px;line-height:120%;color:#020111;width:100%;margin-bottom:21px}.modal__user-photo{width:48px;height:48px;border-radius:50%}.modal__link{display:-webkit-flex;display:flex;-webkit-align-items:center;align-items:center;gap:10px;font-weight:400;font-size:18px;line-height:140%;color:#575d65;margin-bottom:16px}.modal__link:last-child{margin-bottom:0}.modal__log-out{padding:24px}.modal__button{display:-webkit-flex;display:flex;-webkit-align-items:center;align-items:center;gap:11px}.hero{margin:48px 0 104px}.hero__donate-block{padding:24px 40px;background:#fff;border-radius:16px;display:-webkit-flex;display:flex;-webkit-align-items:center;align-items:center}.hero__donate-block.profile{margin:48px 0 64px}.hero__figure-donate{display:-webkit-flex;display:flex;-webkit-align-items:center;align-items:center;gap:24px}.hero__figcaption{font-weight:400;font-size:18px;line-height:140%}.hero__donate-link{padding:16px 56px;background:#13161b;border-radius:8px;color:#fff;font-weight:600;font-size:18px;line-height:140%;height:-webkit-fit-content;height:-moz-fit-content;height:fit-content;margin-left:104px;transition:background .3s}.hero__donate-link:hover{background:#242b33}.hero__donate-link:active{background:#0b0d10}.hero__main{background:#13161b;color:#fff;border-radius:16px;padding:48px 40px}.hero__button{display:block;width:-webkit-fit-content;width:-moz-fit-content;width:fit-content}.hero__title{font-weight:500;font-size:64px;line-height:110%;margin-bottom:16px}.hero__subtitle{font-size:21px;line-height:140%;margin-bottom:88px;color:#c6cad2}.hero__figure{display:-webkit-flex;display:flex;gap:10px;-webkit-align-items:center;align-items:center;color:#24292f}.hero__button:hover figcaption{color:#3f3a92}.hero__button:hover path{fill:#3f3a92}.hero__button:active figcaption{color:#2f2c6d}.hero__button:active path{fill:#2f2c6d}.history{margin-bottom:104px}.history__title{font-size:48px;line-height:120%;margin-bottom:16px}.history__subtitle{line-height:140%;color:#7c7b88;width:590px;margin-bottom:40px}.history__main{background:#fff;border-radius:16px;padding:40px}.history__main-values{width:100%;display:-webkit-flex;display:flex;-webkit-justify-content:space-between;justify-content:space-between;margin-bottom:12px;color:#7c7b88}.history__main-values span:last-child{margin-right:24px}.history__list{height:780px;overflow:auto;width:100%}.history__card{display:-webkit-flex;display:flex;-webkit-align-items:center;align-items:center;padding:24px 0;border-top:1px solid #dddce4;margin-right:24px}.history__user{display:-webkit-flex;display:flex;gap:24px;width:480px;margin-right:24px;-webkit-align-items:flex-start;align-items:flex-start}.history__user-picture{width:48px;height:48px}.history__user-name{font-weight:600;font-size:22px;line-height:120%;color:#24292f;margin-bottom:11px;padding-top:11px}.history__link{display:-webkit-flex;display:flex;-webkit-align-items:center;align-items:center;gap:8px;color:#575d65;font-weight:400;font-size:18px;line-height:140%;margin-bottom:8px}.history__link:last-child{margin-bottom:0}.history__user-github,.history__user-name{width:380px;text-overflow:ellipsis;white-space:nowrap;overflow:hidden}.history__badges{display:-webkit-flex;display:flex;gap:2px;font-size:14px;line-height:140%}.history__badge{display:-webkit-flex;display:flex;color:#fff}.history__badge-title{display:block;padding:2px 6px 3px 8px;background:#000;border:1px solid #e4eaf1;border-radius:4px 0 0 4px}.history__badge-count{display:block;padding:2px 6px 3px 8px;background:#6d96ff;border-width:1px 1px 1px 0;border-style:solid;border-color:#e4eaf1;border-radius:0 4px 4px 0}.history__user-time{margin-left:auto;color:#7c7b88}.history ::-webkit-scrollbar{width:10px}.history ::-webkit-scrollbar-bottom{width:4px}.history ::-webkit-scrollbar-track{background-color:#f2f2f8}.history ::-webkit-scrollbar-thumb{background-color:#b8b6e2;border-radius:10px}.example{margin-bottom:104px}.example__title{font-weight:500;font-size:32px;line-height:120%;margin-bottom:32px;color:#13161b}.example__badges-group{font-size:27px;display:-webkit-flex;display:flex;gap:5px;width:100%;margin-bottom:33px}.example .history__badge{border-radius:8px;overflow:hidden}.example .history__badge-title{padding:4px 12px 6px 16px;border-radius:8px 0 0 8px;margin-right:1px}.example .history__badge-count{padding:4px 12px 6px 16px;border-radius:0 8px 8px 0}.example__title.center{margin-left:312px}.example__chart-group{display:grid;grid-template-columns:1fr 1fr;gap:32px}.instruction__header{margin-bottom:80px}.instruction__step{display:-webkit-flex;display:flex;-webkit-align-items:center;align-items:center;gap:7px;color:#575d65;font-weight:400;font-size:18px;line-height:140%;margin-bottom:16px}.instruction__header-group{display:-webkit-flex;display:flex;-webkit-justify-content:space-between;justify-content:space-between;-webkit-align-items:flex-end;align-items:flex-end}.instruction__title{font-weight:500;font-size:44px;line-height:120%;width:440px}.instruction__show-button{display:-webkit-flex;display:flex;-webkit-align-items:center;align-items:center;gap:12px;padding:16px 16px;font-weight:600;font-size:18px;line-height:140%;background:#13161b;border-radius:8px;color:#fff;transition-property:background,-webkit-transform;transition-property:transform,background;transition-property:transform,background,-webkit-transform;transition-duration:.3s}.instruction__show-button:hover{background:#242b33}.instruction__show-button:active{background:#0b0d10}.instruction__show-img.active{-webkit-transform:rotate(180deg);transform:rotate(180deg)}.instruction__steps{display:-webkit-flex;display:flex;-webkit-flex-direction:column;flex-direction:column;gap:80px;margin-bottom:103px}.instruction__steps.hide{display:none}.header.step-block{height:32px}.header.step-block .header__wrapper{width:623px}.header.step-block .header__logo-img{width:60px;height:19px}.button__log-in.hero-button.step-block{padding:8px 28px}.button__log-in.hero-button.step-block .button__log-in-img{width:12px;height:12px}.button__log-in.hero-button.step-block .button__text{font-size:9px}.button__log-in.hero-button.step-block .button__log-in-link{display:-webkit-flex;display:flex;-webkit-align-items:center;align-items:center;gap:6px}.button__log-in.step-block{padding:6.14671px 20.489px;gap:6.15px;border-radius:4.09781px}.button__log-in.step-block .button__log-in-img{width:12px;height:12px}.button__log-in.step-block .button__text{font-size:11px}.button__log-in.step-block .button__log-in-link{display:-webkit-flex;display:flex;-webkit-align-items:center;align-items:center;gap:6px}.hero.step-block{margin:40px 56px}.hero.step-block .hero__main{padding:24px 20px}.hero.step-block .hero__title{font-size:32px}.hero.step-block .hero__subtitle{font-size:10px;margin-bottom:45px}.example__badges.step-block{margin:0 56px}.example__badges.step-block .example__title{font-size:16px;margin-bottom:20px}.example__badges.step-block .example__badges-group{gap:2px;font-size:14px;margin-bottom:5px}.step{display:-webkit-flex;display:flex;-webkit-justify-content:flex-end;justify-content:flex-end;gap:56px}.step__progress{display:-webkit-flex;display:flex;-webkit-flex-direction:column;flex-direction:column;-webkit-align-items:center;align-items:center;gap:24px;width:-webkit-max-content;width:-moz-max-content;width:max-content}.step__done{width:48px;min-height:48px;border-radius:50%;border:2px solid #575d65}.step__done-img{display:none;width:48px;height:48px;position:relative;top:50%;left:50%;-webkit-transform:translate(-50%,-50%);transform:translate(-50%,-50%)}.step__browser{display:block}.step__view-group{display:-webkit-flex;display:flex;-webkit-flex-direction:column;flex-direction:column;gap:32px}.step__view{width:800px}.step__main{padding:32px 29px;background:#fff;border-radius:16px}.step__progress-bar{width:2px;height:100%;background:#dddce4}.step__title{font-weight:500;font-size:32px;line-height:120%;margin-bottom:32px}.step__subtitle{display:block;font-weight:400;font-size:18px;line-height:140%;color:#6d7280;margin-bottom:32px}.step__title.step-link{margin-bottom:16px}.step__link{font-weight:400;font-size:18px;line-height:140%;color:#0057ff;display:-webkit-flex;display:flex;margin-bottom:32px;-webkit-align-items:center;align-items:center}.step__link:hover{font-weight:500}.step.active .step__done-img{display:block}.step.active .step__progress-bar{background:#2da44e}.wrapper-u8.step-wrapper{background:#f2f2f8;min-width:736px}.step-2__item-group{display:-webkit-flex;display:flex;gap:10px;margin-bottom:16px;font-size:14px}.step-2__select-block{display:-webkit-flex;display:flex;-webkit-align-items:center;align-items:center;padding:9.5px 16px;background:#f6f8fa;border:.5px solid #d9dbdd;border-radius:4px;gap:6px}.step-2__select{cursor:pointer}.step-2__select:active,.step-2__select:focus{border:0;outline:0}.step-2__profile-photo{width:26px;height:26px;border-radius:50%}.step-2__input-block{padding:9.5px 16.5px;background:#f6f8fa;border:.5px solid #d9dbdd;border-radius:4px;display:-webkit-flex;display:flex;-webkit-align-items:center;align-items:center}.step-2__input{border:0;background:inherit}.step-2__input:active,.step-2__input:focus{border:0;outline:0}.step-2__label{display:inline-block;font-weight:500;font-size:12px;line-height:140%;margin-bottom:8px}.step-2__slash{font-weight:400;font-size:20px;line-height:120%;margin-top:auto;margin-bottom:10px}.step-2__label-star{color:red;font-size:10px}.step-2__readme{display:-webkit-flex;display:flex;gap:16px;background:#dff4ff;border:.5px solid #c7d8e2;border-radius:8px;padding:24px 16px;margin-bottom:16px}.step-2__readme-img{width:42px;height:35px}.step-2__readme-text{font-weight:400;font-size:15px;line-height:160%;color:#363343}.step-2__readme-strong{font-weight:600}.step-2__readme-italic{font-style:italic}.step-2__readme-file{padding:16px 0;border-top:.5px solid #d8dee4;border-bottom:.5px solid #d8dee4;margin-bottom:16px}.step-2__checkbox{width:16px;height:16px;background:#0075ff;border-radius:4px}.step-2__create-repo{padding:7px 16px 7px;background:#2da44e;border:1px solid #2a9047;border-radius:6px;transition:border-color .2s,background .2s}.step-2__create-repo:hover{border-color:rgba(27,31,36,.1490196078);background:#2c974b}.step-2__create-link{font-weight:600;font-size:14px;line-height:140%;color:#fff}.step-3__title{font-weight:500;font-size:22px;line-height:120%;margin-bottom:12px}.step-3__item-group{display:grid;grid-template-columns:1fr 1fr;gap:32px}.step-3__img{border-radius:16px}.step-3__item-log{position:relative}.step-3__item{max-width:353px}.step-3__button{width:-webkit-max-content;width:-moz-max-content;width:max-content;position:absolute;top:50%;left:50%;-webkit-transform:translate(-50%,-25%);transform:translate(-50%,-25%)}.step-3__item-content{display:-webkit-flex;display:flex;-webkit-justify-content:space-between;justify-content:space-between;gap:16px;padding:32px;background:#f2f2f8;border-radius:16px}.step-3__item-text{word-wrap:break-word;width:224px}.step-3__copy-text{position:relative;background:#fcfcff;border:1px solid #e5e4f1;border-radius:4px;padding:12px;height:-webkit-fit-content;height:-moz-fit-content;height:fit-content;-webkit-animation-duration:2s;animation-duration:2s;width:48px;height:48px}.step-3__copy-text:hover{background:#f2f2f8}.step-3__copy-img{display:block;pointer-events:none}.step-3__copy-done{display:none;position:absolute;top:50%;left:50%;-webkit-transform:translate(-50%,-50%);transform:translate(-50%,-50%)}.step-3__copy-popup{position:absolute;left:-92px;top:6px;font-weight:500;font-size:18px;line-height:140%;padding:4px 8px;background:#13161b;border-radius:4px;color:#fff}@-webkit-keyframes github-button{0%,to{background:#f2fffb;border:1px solid #1a7b63;box-shadow:0 0 2px rgba(26,123,99,.6),0 0 8px rgba(26,123,99,.4)}}@keyframes github-button{0%,to{background:#f2fffb;border:1px solid #1a7b63;box-shadow:0 0 2px rgba(26,123,99,.6),0 0 8px rgba(26,123,99,.4)}}.step-4{padding:32px 0}.step-4__nav{width:100%;border-bottom:1px solid #d8dee4;padding-left:235px;margin-bottom:15px}.step-4__nav-list{display:-webkit-flex;display:flex;gap:5px}.step-4__nav-link{display:-webkit-flex;display:flex;-webkit-align-items:center;align-items:center;gap:5.5px;padding:0 6px;font-size:8.58333px;line-height:18px;padding-bottom:5.5px;color:#24292f}.step-4__nav-item:first-child{border-bottom:1.4px solid #fd8c73}.step-4__profile{display:-webkit-flex;display:flex;padding:0 25px 32px 29px;gap:15px}.step-4__left-section{margin-top:-40px}.step-4__profile-photo{margin-bottom:8px}.step-4__photo{width:181px;height:181px;border-radius:50%}.step-4__profile-name{padding-left:5px;font-weight:300;font-size:12.2619px;line-height:15px;color:#57606a;margin-bottom:9px}.step-4__edit-profile{width:100%}.step-4__edit-profile-link{color:#24292f;padding:5px 16px;font-size:8px;line-height:12px;border:.61px solid rgba(27,31,36,.15);box-shadow:0 .6px 0 rgba(27,31,36,.04),inset 0 .6px 0 rgba(255,255,255,.25);border-radius:3.6px;background:#f6f8fa;display:block;width:100%;text-align:center}.step-4__edit-profile-link:hover{background:#f3f4f6;border-color:rgba(38,40,43,.1490196078);transition:.3s;transition-property:background,border-color}.step-4__right-section{width:100%}.step-4__readme-file{position:relative;padding:15px 16px 15px 18px;border:.613095px solid #d0d7de;border-radius:3.67857px;width:100%;margin-bottom:15px}.step-4__readme-user{font-weight:400;font-size:7.3px;line-height:11px;margin-bottom:8px}.step-4__readme-code{font-weight:600;font-size:10.7px;line-height:13px}.step-4__edit{position:absolute;top:15px;right:14px}.step-4__popular-repo{max-width:245px}.step-4__popular-title{font-weight:400;font-size:9.80952px;line-height:15px;margin-bottom:5px}.step-4__popular-item{padding:10px 12px 25px;border:.613095px solid #d0d7de;border-radius:3.67857px}.step-4__popular-item{display:-webkit-flex;display:flex;-webkit-justify-content:space-between;justify-content:space-between}.step-4__repo-link{font-weight:600;font-size:8.5px;line-height:13px;color:#0969da}.step-4__repo-setings{font-size:7.35px;padding:.6px 6px .6px 6px;display:-webkit-flex;display:flex;-webkit-align-items:center;align-items:center;color:#57606a;border:.61px solid #d0d7de;border-radius:14.7px}.step-4__readme-description{font-size:9.8px;line-height:15px;padding:9px 14px;background:#dafbe1;border:.6125px solid #a0e5b2;border-radius:4.9px;margin-bottom:10px}.step-4__file-user-name{font-weight:500}.step-4__file{border-radius:3.675px;border:.6125px solid #d0d7de;overflow:hidden}.step-4__file-buttons{display:-webkit-flex;display:flex;background:#f6f8fa;border-bottom:.61px solid #d0d7de}.step-4__file-edit{font-size:9.8px;line-height:15px;display:-webkit-flex;display:flex;-webkit-align-items:center;align-items:center;gap:5px;padding:5px 10px;border-width:0 .6125px 0 0;border-style:solid;border-color:#d0d7de;border-radius:3.675px 3.675px 0 0;background-color:#fff;margin-bottom:-1px}.step-4__file-preview{font-size:9.8px;line-height:15px;display:-webkit-flex;display:flex;-webkit-align-items:center;align-items:center;gap:5px;padding:5px 10px;background:#f6f8fa;position:relative;z-index:2}.step-4__code{padding:14px 10px 35px 10px;font-weight:500;font-size:8.575px;line-height:160%}.step-4__code-list{list-style:auto;color:#6e7781;padding-left:10px}.step-4__code-item{padding-left:15px}.step__main-padding{padding:32px 29px}.step-5{background:#6d96ff;display:-webkit-flex;display:flex;-webkit-flex-direction:column;flex-direction:column;gap:16px}.stat-reg{margin-top:64px}.stat-reg__title{font-weight:500;font-size:44px;line-height:120%;margin-bottom:32px}.profile{margin-top:64px;color:#111028}.profile__card{display:-webkit-flex;display:flex;-webkit-align-items:center;align-items:center;gap:32px;margin-bottom:40px}.profile__photo{width:128px;height:128px;border-radius:50%}.profile__info{display:-webkit-flex;display:flex;-webkit-flex-direction:column;flex-direction:column;gap:8px}.profile__name{font-weight:500;font-size:32px;line-height:120%;color:#020111}.profile .history__badges{margin-left:auto}.profile__link{display:block;font-size:18px;line-height:140%;color:#575d65;display:-webkit-flex;display:flex;-webkit-align-items:center;align-items:center;gap:8px}.profile__statistics{background:#fff;border-radius:16px;width:100%;height:488px;padding:40px;margin-bottom:104px;color:#111028}.profile__header{display:-webkit-flex;display:flex;-webkit-align-items:flex-end;align-items:flex-end;margin-bottom:40px}.profile__title{font-weight:500;font-size:22px;line-height:120%}.profile__view-count{font-weight:400;font-size:18px;line-height:140%;margin-left:auto;display:-webkit-flex;display:flex;-webkit-align-items:center;align-items:center;color:#575d65}.profile__circle{display:inline-block;width:16px;height:16px;border-radius:50%;background:#6d96ff;margin-right:12px}.profile__select{margin-left:64px;font-weight:500;font-size:18px;line-height:140%}.badge{margin-bottom:104px}.badge__title{font-weight:500;font-size:44px;line-height:120%;margin-bottom:8px}.badge__subtitle{line-height:140%;color:#7c7b88;margin-bottom:40px}.badge__copy{position:relative;margin-left:auto;overflow:hidden}.badge .step__view{width:100%}.badge .step-3__item{max-width:100%}.badge .step-3__item-text{width:100%}.badge .step-3__item-content{padding:32px;background:#fff;border-radius:16px;min-height:189px;font-weight:400;font-size:18px;line-height:140%}.referal__title{font-weight:500;font-size:44px;line-height:120%;margin-bottom:32px}.referal .profile__circle{background:#000}.action{display:-webkit-flex;display:flex;-webkit-align-items:flex-end;align-items:flex-end;-webkit-justify-content:space-between;justify-content:space-between;background:#13161b;border-radius:16px;padding:32px 32px 40px;margin-bottom:104px}.action__title{font-weight:500;font-size:44px;line-height:120%;color:#fff;margin-bottom:16px}.action__subtitle{font-weight:400;font-size:18px;line-height:140%;color:#c6cad2}.action.white{background:#fff}.action.white .action__title{color:#13161b}.action.white .action__subtitle{color:#575d65}.button__log-in{border:1px solid #e4eaf1;border-radius:8px;display:-webkit-flex;display:flex;-webkit-align-items:center;align-items:center;gap:13px;background:#fff;transition:background .3s}.button__log-in:hover{background:#f3f5f8}.button__log-in.black{background:#13161b}.button__log-in.black:hover{background:#242b33}.button__log-in.black:active{background:#0b0d10}.button__log-in.black .button__text{color:#fff}.button__log-in-img{width:24px;height:24px;border-radius:50%}.button__log-in.hero-button .button__link{padding:16.5px 56px}.button__link{display:-webkit-flex;display:flex;-webkit-align-items:center;align-items:center;gap:12px;padding:12px 40px}.button__text{color:#13161b;font-weight:600;font-size:18px}.wrapper-u8{max-width:1216px;margin:0 auto;color:#020111}body{font-size:18px}main{background:#f2f2f8}.sorting{margin:40px 0 24px;padding-bottom:16px;border-bottom:1px solid #dddce4;font-size:18px;display:-webkit-flex;display:flex}.sorting__item{display:-webkit-flex;display:flex;margin-right:32px;position:relative}.sorting__item h4{display:-webkit-flex;display:flex;-webkit-align-items:center;align-items:center}.sorting__item:last-child{margin-right:0}.sorting__title{padding:8px;padding-right:12px}.sorting__select{position:relative;display:-webkit-flex;display:flex;-webkit-align-items:center;align-items:center}.sorting__current{width:-webkit-max-content;width:-moz-max-content;width:max-content;border:0;font-weight:600;font-size:18px;line-height:140%;color:#111028;padding:0;padding-right:8px;text-align:left;background:0 0;position:relative}.sorting__icon{padding:7px 4px}.sorting__list{opacity:0;pointer-events:none;display:-webkit-flex;display:flex;-webkit-flex-direction:column;flex-direction:column;padding:8px 0;border:1px solid #ededf8;box-shadow:0 6px 16px rgba(17,16,40,.15);border-radius:4px;background:#fff;color:#3c3a59;font-size:18px;line-height:140%;position:absolute;bottom:-156px;left:80px;z-index:1;width:-webkit-max-content;width:-moz-max-content;width:max-content}.sorting__option{padding:8px 16px;text-align:left;background:inherit;transition:background .3s;color:#7c7b88}.sorting__option:hover{background:#f2f2f8}.sorting__list.is-visible{transition:opacity .4s;opacity:1;pointer-events:visible}</style>
</head>
<body>
`)
	streamheader(qw422016, ProfileView{}, sessionProfile)
	qw422016.N().S(`
<main class="main">
    <div class="wrapper-u8">
        <section class="hero">
            `)
	if showCharity {
		qw422016.N().S(`
            <div class="hero__donate-block">
                `)
		streamcharity(qw422016)
		qw422016.N().S(`
            </div>
            `)
	}
	qw422016.N().S(`
            <div class="hero__main">
                <h1 class="hero__title">Track your GitHub profile views</h1>
                <p class="hero__subtitle">
                    Receive, view and analyze your profile views and profile performance statistics
                </p>
                <button class="button__log-in hero-button">
                    `)
	if sessionProfile.ID == 0 {
		qw422016.N().S(`
                    <a href="/login/github" class="button__link">
                        <img
                                src="/assets/images/github.svg"
                                width="24"
                                height="24"
                                alt="github"
                                class="button__log-in-img"
                        />
                        <span class="button__text">Log in with GitHub</span>
                    </a>
                    `)
	} else {
		qw422016.N().S(`
                    <a href="https://u8views.com/github/`)
		qw422016.E().S(sessionProfile.Username)
		qw422016.N().S(`" class="button__link">
                        <img
                                src="https://avatars.githubusercontent.com/u/`)
		qw422016.E().S(sessionProfile.SocialProviderUserID)
		qw422016.N().S(`?v=4&s=48"
                                width="24"
                                height="24"
                                alt="`)
		qw422016.E().S(sessionProfile.GetName())
		qw422016.N().S(` profile photo"
                                class="button__log-in-img"
                        />
                        <span class="button__text">View your statistics</span>
                    </a>
                    `)
	}
	qw422016.N().S(`
                </button>
            </div>
        </section>
    </div>
</main>
`)
	streamfooter(qw422016)
	qw422016.N().S(`
</body>
</html>
`)
}

func WriteIndex(qq422016 qtio422016.Writer, sessionProfile ProfileView, showCharity bool, registrationHistoryProfiles []FullProfileView) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	StreamIndex(qw422016, sessionProfile, showCharity, registrationHistoryProfiles)
	qt422016.ReleaseWriter(qw422016)
}

func Index(sessionProfile ProfileView, showCharity bool, registrationHistoryProfiles []FullProfileView) string {
	qb422016 := qt422016.AcquireByteBuffer()
	WriteIndex(qb422016, sessionProfile, showCharity, registrationHistoryProfiles)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}
