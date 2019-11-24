function highlightText(element, text){
  var innerHTML = element.innerHTML;
  var index = innerHTML.indexOf(text);
  if (index >= 0) { 
    innerHTML = innerHTML.substring(0,index) + "<span class='highlight'>" + innerHTML.substring(index,index+text.length) + "</span>" + innerHTML.substring(index + text.length);
    element.innerHTML = innerHTML;
   }
}

function scrollAndHighlight(search_str, n){
  var currentNumber = 0;
  var BreakException;
  var nodes = get_nodes_containing_text("*", search_string);
  try{
    nodes.forEach(element=>{
      if(currentNumber == parseInt(n)){
        highlightText(element, search_str);
        element.scrollIntoView({block: "start", behavior: "smooth"});
        throw BreakException;
      };
      currentNumber++;
    })
  }catch(e){
    if (e != BreakException) throw e;
  }
}

function get_nodes_containing_text(selector, text) {
  const elements = [...document.querySelectorAll(selector)];

  return elements.filter(
    (element) =>
      element.childNodes[0]
      && element.childNodes[0].nodeValue
      && RegExp(text, "u").test(element.childNodes[0].nodeValue.trim())
  );
}

function getAllUrlParams(url) {

    // get query string from url (optional) or window
    var queryString = url ? url.split('?')[1] : window.location.search.slice(1);
  
    // we'll store the parameters here
    var obj = {};
  
    // if query string exists
    if (queryString) {
  
      // stuff after # is not part of query string, so get rid of it
      queryString = queryString.split('#')[0];
  
      // split our query string into its component parts
      var arr = queryString.split('&');
  
      for (var i = 0; i < arr.length; i++) {
        // separate the keys and the values
        var a = arr[i].split('=');
  
        // set parameter name and value (use 'true' if empty)
        var paramName = a[0];
        var paramValue = typeof (a[1]) === 'undefined' ? true : a[1];
  
        // (optional) keep case consistent
        paramName = paramName.toLowerCase();
        if (typeof paramValue === 'string') paramValue = paramValue.toLowerCase();
  
        // if the paramName ends with square brackets, e.g. colors[] or colors[2]
        if (paramName.match(/\[(\d+)?\]$/)) {
  
          // create key if it doesn't exist
          var key = paramName.replace(/\[(\d+)?\]/, '');
          if (!obj[key]) obj[key] = [];
  
          // if it's an indexed array e.g. colors[2]
          if (paramName.match(/\[\d+\]$/)) {
            // get the index value and add the entry at the appropriate position
            var index = /\[(\d+)\]/.exec(paramName)[1];
            obj[key][index] = paramValue;
          } else {
            // otherwise add the value to the end of the array
            obj[key].push(paramValue);
          }
        } else {
          // we're dealing with a string
          if (!obj[paramName]) {
            // if it doesn't exist, create property
            obj[paramName] = paramValue;
          } else if (obj[paramName] && typeof obj[paramName] === 'string'){
            // if property does exist and it's a string, convert it to an array
            obj[paramName] = [obj[paramName]];
            obj[paramName].push(paramValue);
          } else {
            // otherwise add the property
            obj[paramName].push(paramValue);
          }
        }
      }
    }
  
    return obj;
  }

var url = window.location.href
console.log(url)
var search_string = getAllUrlParams(url).search_string
var entry_number = getAllUrlParams(url).n

// we got search string and entry number, so we should highlight and scroll
if (search_string != undefined){
  console.log(search_string +":"+entry_number)
  
  scrollAndHighlight(search_string, entry_number)
}
// scroll and highlight

