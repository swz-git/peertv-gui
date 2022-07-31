export namespace main {
	
	export class SearchResult {
	    url: string;
	    title: string;
	    seed: string;
	    leech: string;
	    magnetlink?: string;
	    filesize: number;
	    engine: string;
	    score: number;
	
	    static createFrom(source: any = {}) {
	        return new SearchResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.url = source["url"];
	        this.title = source["title"];
	        this.seed = source["seed"];
	        this.leech = source["leech"];
	        this.magnetlink = source["magnetlink"];
	        this.filesize = source["filesize"];
	        this.engine = source["engine"];
	        this.score = source["score"];
	    }
	}

}

