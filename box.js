class GridBox extends HTMLElement {
    value = 0;
    
    constructor() {
        super();
        this.innerHTML = "I am a box"
    }
}

customElements.define("grid-box", GridBox);

class Grid extends HTMLElement {
    constructor() {
        super()
        this.innerHTML = "<grid-box></grid-box><grid-box></grid-box><grid-box></grid-box><grid-box></grid-box>"
    }
}
customElements.define("grid-element", Grid);
