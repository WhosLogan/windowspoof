export namespace main {
	
	export class Window {
	    pid: number;
	    name: string;
	    windowName: string;
	
	    static createFrom(source: any = {}) {
	        return new Window(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.pid = source["pid"];
	        this.name = source["name"];
	        this.windowName = source["windowName"];
	    }
	}

}

