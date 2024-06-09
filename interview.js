// in nodejs it will not be executed and it says blockScope is not defined
// but in browser it will be executed and it will print "something cool...."
// so in order to prevent this we can use "use strict" in the beginning of the file and it will throw an error in browser as well.

{
  function blockScope() {
    console.log("something cool....");
    return "coool stuff...";
  }
}

// blockScope();

function fizz() {
  this.name = "fizz";
  console.log(this.name);
}

const fi = new fizz();
fi.name = "dusty....";
console.log(fi);
