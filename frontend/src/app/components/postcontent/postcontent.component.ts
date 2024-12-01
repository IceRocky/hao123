import { Component } from '@angular/core';
import { feedTitles } from 'src/app/data/feed-title';
import { TopNewsService } from 'src/app/services/top-news.service';
@Component({
  selector: 'app-postcontent',
  templateUrl: './postcontent.component.html',
  styleUrls: ['./postcontent.component.css']
})
export class PostcontentComponent {
  feedTitles: any[] = feedTitles;
  topNews: any[] = [];
  constructor( private TopNewsService : TopNewsService) {}
  ngOnInit(): void {
    this.fetchTopNews(1, 10);
  }
  fetchTopNews(pageId: number, pageSize: number): void {
    this.TopNewsService.getTopNews(pageId, pageSize).subscribe(
      (response: any) => {
        this.topNews = response.topNews;
      },
      (error: any) => {
        console.error(error);
      }
    );
  }


}
