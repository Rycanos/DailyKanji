export namespace character {
	
	export class Character {
	    CharId: number;
	    CharStroke: number;
	    JlptLvl: number;
	    Char: string;
	    ReadingJoyo: string;
	    MeaningOn: string;
	    MeaningKun: string;
	
	    static createFrom(source: any = {}) {
	        return new Character(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.CharId = source["CharId"];
	        this.CharStroke = source["CharStroke"];
	        this.JlptLvl = source["JlptLvl"];
	        this.Char = source["Char"];
	        this.ReadingJoyo = source["ReadingJoyo"];
	        this.MeaningOn = source["MeaningOn"];
	        this.MeaningKun = source["MeaningKun"];
	    }
	}

}

