function valid(){
    v = document.getElementsByName("search_key")[0].value;

    console.log(v)
    if ("" == v){
        alert("请填写搜索关键词")
        return false;
    }

    return true;
}