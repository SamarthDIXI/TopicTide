let producerSocket=new WebSocket('ws://localhost:8080');
producerSocket.onopen=()=>{
    console.log("Producer-Broker connection opened!");
}
producerSocket.onerror=()=>{
    console.error("Producer-Broker websocket error: ",error)
}
function createJSON(event){
    event.preventDefault()
    let topic=document.getElementById("i1").value;
    let content=document.getElementById("i2").value;
    let producerMsg={
        "Topic":topic,
        "Content":content
    };
    console.log(JSON.stringify(producerMsg));
    producerSocket.send(JSON.stringify(producerMsg));
}
document.getElementById("f1").addEventListener("submit",(event)=>{
    createJSON(event);
});