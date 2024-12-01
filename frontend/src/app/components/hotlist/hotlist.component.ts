import { Component } from '@angular/core';
import { HotListService } from 'src/app/services/hot-list.service';

@Component({
  selector: 'app-hotlist',
  templateUrl: './hotlist.component.html',
  styleUrls: ['./hotlist.component.css']
})
export class HotlistComponent {
  hotLists: any[] = [];
  expandedIndex: number | null = 0;
  constructor(    private HotListService : HotListService,) {

  }
  ngOnInit(): void {
    this.fetchHotList(1, 10);
  }


  // hotList
  fetchHotList(pageId: number, pageSize : number) :void{
    this.HotListService.gethotList(pageId, pageSize).subscribe(
      (data) => {
        this.hotLists = data.hotLists;
      },
      (error) =>{
        console.error('Error fetching HotLists:', error);
      }
    )
  };

}
