import { CollectionViewer, DataSource } from "@angular/cdk/collections";
import { BehaviorSubject, Observable, of, Subscription } from 'rxjs';
import { catchError } from 'rxjs/internal/operators/catchError';
import { NewsItemModel } from '../../../data/newsitem.model';
import { NewsService } from "../../services/news.service";

export class NewsDataSource extends DataSource<NewsItemModel> {

    private cache = Array.from<NewsItemModel>({ length: 0 });
    private newsSubject = new BehaviorSubject<NewsItemModel[]>(this.cache);
    
    private subscription = new Subscription();

    private pageSize = 200;
    private lastPage = 0;

    private searchID = "";

    constructor(private newsService: NewsService) {
        super();
    }

    connect(collectionViewer: CollectionViewer): Observable<NewsItemModel[]> {
        this.subscription.add(collectionViewer.viewChange.subscribe(range => {
            if (range.end >= this.cache.length) {
                this.loadData("");
            }
        
          }));
          return this.newsSubject;
    }

    disconnect(collectionViewer: CollectionViewer): void {
        this.newsSubject.complete();
        this.subscription.unsubscribe();
    }

    setSearchId(searchID: string) {
        this.searchID = searchID;
        this.cache.length = 0;
        this.loadData(searchID);
    }

    loadData(searchID: string) {
        let key = "";
        let sortkey = null;

        if (this.cache.length > 0) {
            key = this.cache[this.cache.length - 1].id;
            sortkey = this.cache[this.cache.length - 1].datetime;
        }

        this.newsService.getNewsItems(200, key, sortkey, this.searchID).pipe(
            catchError(() => of([]))
        )
        .subscribe(data => {
            this.cache = this.cache.concat(data);
            this.newsSubject.next(this.cache);
        });
    }  

}