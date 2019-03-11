# 1. How to ues RESTfil API Doc.  
- The API doc is in the Postman. You alredy know the address of postman.  
- Yo u can get the RESTful API address as selecting Environment like below.  
    - Use QA or Live(not setting yet)  
    ![image#1]  

- You can see request example and response example like below.  
    ![image#2] 

# 2. How to auth / start Game through Web  
1. User login on Web page  
2. Web server save the JWT from platform RESTful API into the cookie when login  
3. User click the game start button on web page.  
4. Web page run the game launcher. (Naddic will give this to you)  
    - Web page or web server give the JWT in the cookie to the game launcher as parameter.  
5. The game launcher send the JWT to platform RESTful API for validation and getting user information.  
6. The game started.  

# 3. How to development billing (charging virtual money.)
### Charging process  
1. User login on Web page. And then move to charge web page  
2. Web server call the platform RESTful API to get charge item list.  
    + {{url}}/v1/billing/getChargeItems  

3. User see charge item list and click buy button of an item.  
4. Web server call the platform RESTful API for getting payment xsolla token.  
    1. {{url}}v1/billing/getPaymentToken  
    2. The response data like that when success  

    ![image#3]  

5. Set javascript value with mode and token and then Xsolla’s payment gateway UI screen will be opened automatically.  
    1. You see number #4 of this doc for developing charging.  
    i. develop & apply payment gate way Xsolla (Thailand) on the web  

6. If user succeed the charge payment, redirect to a web page like my coin charge history.  
    1. For getting coin charge history. Call the RESTful API  
        + {{url}}/v1/billing/getChargeHistory  

# 4. develop & apply payment gateway Xsolla (Thailand) on the web  
1. Include xsolla java script into your web page.  
```` <Script url="https://static.xsolla.com/embed/paystation/1.0.7/widget.min.js"/> ````

2. see below a example. This is exmaple how to open Xsolla(Thailand) payment gateway popup when user click buy item button.  
```` 
<script>
    var options = {
        access_token: 'ACCESS_TOKEN', //TODO use access token, received on previous step
        sandbox: true //TODO please do not forget to remove this setting when going live
    };
    var s = document.createElement('script');
    s.type = "text/javascript";
    s.async = true;
    s.src = "https://static.xsolla.com/embed/paystation/1.0.7/widget.min.js";
    s.addEventListener('load', function (e) {
        XPayStationWidget.init(options);
    }, false);
    var head = document.getElementsByTagName('head')[0];
    head.appendChild(s);
</script>
<button data-xpaystation-widget-open>Buy Credits</button>
```` 

* access_token, sandbox
1. You can get these from platform RESTful API  
    - {{url}}v1/billing/getPaymentToken  

2. Response value “mode” from platform RESTful API is the sandbox value. So if you get “mode”: “sandbox”, you should set sandbox: true. If you get “mode”:”prod”, you should set sandbox: false.  

3. Explained above #3.(How to development billing (charging virtual money.)).  

4. More detail (xsolla web site):  
https://github.com/xsolla/paystation-embed   
https://developers.xsolla.com/doc/pay-station/  

5. This is my example source code.You can see my test source with react and mobx.  
```` 
async openPay(UID, itemid, history, appState) {
    var complete = false;
    var waiting = false;
    var failed = false;
    var invoice;

    var success_url = 'http://localhost:3000/payment/history'

    var s = document.createElement('script');
    s.type = 'text/javascript';
    s.async = true;
    s.src = '//static.xsolla.com/embed/paystation/1.0.7/widget.min.js';

    s.addEventListener(
        'load',
        function (e) {
            XPayStationWidget.on('close', function () {
                if (complete) history.push('/payment/history'); 
                if (failed) window.location.href = '';
            });

            XPayStationWidget.on(
                'status-invoice status-delivering status-troubled status-done',
                function (event, data) {

                    if (arguments[0].type == 'status-done') {
                        complete = true;
                    }

                    if (arguments[0].type == 'status-delivering')
                        waiting = true;
                    if (arguments[0].type == 'status-troubled') failed = true;
                }
            );
        },
        false
    );

    var head = document.getElementsByTagName('head')[0];
    head.appendChild(s);

    //check inputs
    if ((!UID) || (!itemid)) {
        //go to login
        history.push('/login');
    } else {
        var options = {
            access_token: '',
            lightbox: {
                spinnerColor: '#cccccc',
                closeByClick: false,
                closeByKeyboard: false
            },
            sandbox: false
        };

        // check cookie
        let cookieInfo = null;
        cookieInfo = storage.get('___GOM___'); //cookie name for test

        if (cookieInfo) {
            let paymentToken = null;
            try {
                paymentToken = await BillingAPI.getPaymentToken({
                    token: cookieInfo.token,
                    itemid
                });
            } catch (err) {
                this.setErrorFlashMessage(err.response.data.message);
            }

            appState.setLoading('off');

            if (!paymentToken) {
                this.this.setErrorFlashMessage('something wrong. please try again.');
            } else {
                options.access_token = paymentToken.data.data.token;

                if (paymentToken.data.data.mode === 'sandbox') {
                    options.sandbox = true;
                }

                XPayStationWidget.init(options);
                document.getElementById('buyXsolla').click();
            }
        }else{
            console.log('need login first.');
            await appState.setInitLoggedInUserInfo();
            this.setErrorFlashMessage('need login first.');
            history.push('/login');
        }
    }
}
````  






    
    

