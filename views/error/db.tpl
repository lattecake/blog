<style type="text/css">
    body {
        margin: 0px;
        padding: 0px;
        font-family: "微软雅黑", Arial, "Trebuchet MS", Verdana, Georgia, Baskerville, Palatino, Times;
        font-size: 16px;
    }

    div {
        margin-left: auto;
        margin-right: auto;
    }

    a {
        text-decoration: none;
        color: #1064A0;
    }

    a:hover {
        color: #0078D2;
    }

    img {
        border: none;
    }

    h1, h2, h3, h4 {
        /*	display:block;*/
        margin: 0;
        font-weight: normal;
        font-family: "微软雅黑", Arial, "Trebuchet MS", Helvetica, Verdana;
    }

    h1 {
        font-size: 44px;
        color: #0188DE;
        padding: 20px 0px 10px 0px;
    }

    h2 {
        color: #0188DE;
        font-size: 16px;
        padding: 10px 0px 40px 0px;
    }

    #page {
        width: 910px;
        padding: 20px 20px 40px 20px;
        margin-top: 80px;
    }

    .button {
        width: 180px;
        height: 28px;
        margin-left: 0px;
        margin-top: 10px;
        background: #009CFF;
        border-bottom: 4px solid #0188DE;
        text-align: center;
    }

    .button a {
        width: 180px;
        height: 28px;
        display: block;
        font-size: 14px;
        color: #fff;
    }

    .button a:hover {
        background: #5BBFFF;
    }
</style>

<div id="page" style="border-style:dashed;border-color:#e4e4e4;line-height:30px;">
    <h1>{{.errorCode}}~</h1>
    <h2>Sorry, the site now can not be accessed. </h2>
    <font color="#666666">{{.message}}</font><br/><br/>
    <div class="button">
        <a href="https://lattecake.com/" title="Home" target="_blank">Home</a>
    </div>
</div>
