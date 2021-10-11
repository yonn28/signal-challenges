function areSimilar(a, b) {
    if(a.toString() == b.toString()){
        return true;
    }
    let c = []
    let d = []
    a.map((value,idx)=>{
        if(value != b[idx]){
            c.push(value);
            d.push(b[idx]);
        }
    })

    d = d.reverse();
    if(c.length ===2 && (c.toString() == d.toString())){
        return true
    }
    return false;
}

console.log(areSimilar([1, 2, 3],[2, 1, 3]))