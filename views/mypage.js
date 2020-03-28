///<reference path="./jquery-3.3.1.min.js"/>

/**
 * @typedef {Object} Cresult
 * @property {String} Country
 * @property {String} Cases
 * @property {String} Deaths
 * @property {String} Recovered
 */

function createCountrySelector(){
    var doc = $('#countryselector');
    $.get("/ajax?atype=getlistcountry",function(r){
        var obj = JSON.parse(r);

        /** @type {string[]} */
        var list = obj.Countries
        list.sort();

        $.each(list,function(x,ctr){ 

            var opt = $('<option></option>').appendTo(doc)
            opt.text(ctr);
            if(ctr == obj.Selected){
                opt.attr("selected","selected");
            }
            
        })

        showBySelected()
    })
}


function showBySelected(){
    
    var selected = $('#countryselector').find('option:selected').text();
    showResult(selected)    
    if(!intervalstarted){
        intervalstarted = true;

        setInterval(function(){
            showBySelected(); 
        },1000)
    }
}


 /** @type {Cresult} */
var savedJson = null 
var intervalstarted = false;

function indicBlink(stateblink){
    var d = $('#loopindic')
    d.css({
        backgroundColor : "#ff0000"
    })

    setTimeout(function(){
        d.css({
            backgroundColor : "#ffffff"
        })
    },500) 
}

function showResult(country){
    
    indicBlink();
    $.get("/ajax?atype=getbycountry&country=" + country,function(r){
        /** @type {Cresult} */
        var obj = JSON.parse(r);
        $('.bcases').text(obj.Cases)
        $('.bdeath').text(obj.Deaths)
        $('.brec').text(obj.Recovered)
        $('.scountry').text(obj.Country) 
        if(savedJson != null && savedJson.Country == obj.Country && (
            savedJson.Cases != obj.Cases ||
            savedJson.Deaths != obj.Deaths ||
            savedJson.Recovered != obj.Recovered
        )){             
            playNotif()
        }


        savedJson = obj
    })

}

function playNotif(){
    $('audio')[0].play();
}

function openBrowser(url){
    $.ajax({
        url : "/ajax",
        method : "GET",
        data : {
            "atype" : "openbrowser",
            "url" : url,
        },
        success : function(r){

        }
    })
}

$(document).ready(function(){
    createCountrySelector()
})