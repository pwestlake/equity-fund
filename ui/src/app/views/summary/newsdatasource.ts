import { CollectionViewer, DataSource } from "@angular/cdk/collections";
import { BehaviorSubject, Observable, of, Subscription } from 'rxjs';
import { catchError } from 'rxjs/internal/operators/catchError';
import { NewsItemModel } from '../../../data/newsitem.model';
import { NewsItemViewModel } from '../../../data/newsitem.model';
import { NewsService } from "../../services/news.service";
import { map } from 'rxjs/operators';

export class NewsDataSource extends DataSource<NewsItemViewModel> {

    private cache = Array.from<NewsItemViewModel>({ length: 0 });
    private newsSubject = new BehaviorSubject<NewsItemViewModel[]>(this.cache);
    
    private subscription = new Subscription();

    private pageSize = 200;
    private lastPage = 0;

    private searchID = "";

    constructor(private newsService: NewsService) {
        super();
    }

    connect(collectionViewer: CollectionViewer): Observable<NewsItemViewModel[]> {
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
            const viewData: NewsItemViewModel[] = (data as NewsItemModel[]).map((item) => <NewsItemViewModel> {
                catalogref: item.catalogref,
                selected: false,
                id: item.id,
                companycode: item.companycode,
                companyname: item.companyname,
                content: item.content,
                datetime: item.datetime,
                sentiment: item.sentiment,
                title: item.title
            });
            this.cache = this.cache.concat(viewData);
            this.newsSubject.next(this.cache);
        });
    }  

    getItemAtIndex(index: number): NewsItemModel {
        const item = this.cache[index];
        return item;
    }

    toggleItemSelection(index: number) {
        this.cache[index].selected = !this.cache[index].selected
        if (this.cache[index].selected) {
            // Ensure single selection
            this.cache.forEach((v, i) => {
                if (v.selected && i != index) {
                    v.selected = false;
                }
            });
        }
    }
}