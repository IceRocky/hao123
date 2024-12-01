import { Component } from '@angular/core';
import { CategoryService } from '../services/category.service';
import { SiteItemsService } from '../services/site-items.service';
import { CommonSitesService } from '../services/common-sites.service';
import { GovSitesService } from '../services/gov-sites.service';
import { HotListService } from '../services/hot-list.service';
import { HotListItemService } from '../services/hot-list-item.service';
import { HotSearchsService } from '../services/hot-searchs.service';
import { ItemCategoriesService } from '../services/item-categories.service';
import { ItemsService } from '../services/items.service';
import { TopListsService } from '../services/top-lists.service';
import { TopListItemsService } from '../services/top-list-items.service';
import { HotTabsService } from '../services/hot-tabs.service';
import { feedTitles } from '../data/feed-title';


@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.css']
})
export class DashboardComponent {
  isLongContentHidden = true;

  categories: any[] = [];
  siteItems: any[] = [];
  commonSiteItems: any[] = [];
  govSites: any[] = [];
  hotListItems: any[] = [];
  TopListItems: any[] = [];

  feedTitles: any[] = feedTitles;



  constructor(private categoryService: CategoryService,
    private siteItemService : SiteItemsService,
    private commonSitesService : CommonSitesService,
    private GovSitesService : GovSitesService,

    private HotListItemService : HotListItemService,
    private HotSearchsService : HotSearchsService,
    private ItemCategoriesService : ItemCategoriesService,
    private ItemsService : ItemsService,
    private TopListsService : TopListsService,
    private TopListItemsService : TopListItemsService,
    private HotTabsService : HotTabsService

  ) {}
  ngOnInit(): void {
    this.fetchCategories(1, 10);
    this.fetchsiteItems(1, 10);
    this.fetchCommonSiteItem(1, 10);
    this.fetchGovSites(1, 10);
    this.fetchHotListItem(1,10);

    this.fetchHotSearch(1,10);
    this.fetchItemCategories(1,10);
    this.fetchItems(1,10);
    this.fetchTopLists(1,10);
    this.fetchTopListItems(1,10);
    this.fetchHotTabs(1,10);

  }
 // Togogle show/off long content
  toggleLongContent(): void {
    this.isLongContentHidden = !this.isLongContentHidden;
  }


  // Categories
  fetchCategories(pageId: number, pageSize: number): void {
    this.categoryService.getCategories(pageId, pageSize).subscribe(
      (data) => {
        this.categories = data.categories;
        console.log("asdfsadfsdfsdfasd")
      },
      (error) => {
        console.error('Error fetching categories:', error);
      }
    );
  }

  setFeedTitles  = () => {
    this.feedTitles = feedTitles;
  };

  //  siteItems
    fetchsiteItems(pageId: number, pageSize : number) :void{
      this.siteItemService.getsiteItems(pageId, pageSize).subscribe(
        (data) => {
          this.siteItems = data.siteItems;
        },
        (error) =>{
          console.error('Error fetching siteItems:', error);
        }
      )
    };

    // commonSiteItem
    fetchCommonSiteItem(pageId: number, pageSize : number) :void{
      this.commonSitesService.getCommonSiteItem(pageId, pageSize).subscribe(
        (data) => {
          this.commonSiteItems = data.commonSite;
        },
        (error) =>{
          console.error('Error fetching commonSiteItem:', error);
        }
      )
    };

  // govSites
  fetchGovSites(pageId: number, pageSize : number) :void{
    this.GovSitesService.getGovSiteItem(pageId, pageSize).subscribe(
      (data) => {
        this.govSites = data.govSites;
        console.log(this.govSites+"asdfsadfsadfsadfsd")
      },
      (error) =>{
        console.error('Error fetching GovSites:', error);
      }
    )
  };


  // hotListItem
  fetchHotListItem(pageId: number, pageSize : number) :void{
    this.HotListItemService.gethotListItem(pageId, pageSize).subscribe(
      (data) => {
        this.hotListItems =  data;
        console.log(JSON.stringify(data))
      },
      (error) =>{
        console.error('Error fetching HotLists:', error);
      }
    )
  };

  // hotSearchs
  fetchHotSearch(pageId: number, pageSize : number) :void{
    this.HotSearchsService.getHotSearchs(pageId, pageSize).subscribe(
      (data) => {
        this.hotListItems =  data;
        console.log(JSON.stringify(data))
      },
      (error) =>{
        console.error('Error fetching HotLists:', error);
      }
    )
  };
  //ItemCategoriesService
  fetchItemCategories(pageId: number, pageSize : number) :void{
    this.ItemCategoriesService.getItemCategories(pageId, pageSize).subscribe(
      (data) => {
        this.hotListItems =  data;
        console.log(JSON.stringify(data))
      },
      (error) =>{
        console.error('Error fetching HotLists:', error);
      }
    )
  };
  //ItemsService
  fetchItems(pageId: number, pageSize : number) :void{
    this.ItemsService.getItems(pageId, pageSize).subscribe(
      (data) => {
        this.hotListItems =  data;
        console.log(JSON.stringify(data))
      },
      (error) =>{
        console.error('Error fetching HotLists:', error);
      }
    )
  };
  //TopListsService
  fetchTopLists(pageId: number, pageSize : number) :void{
    this.TopListsService.getTopLists(pageId, pageSize).subscribe(
      (data) => {
        this.hotListItems =  data;
        console.log(JSON.stringify(data))
      },
      (error) =>{
        console.error('Error fetching HotLists:', error);
      }
    )
  };
  //TopListItemsService
  fetchTopListItems(pageId: number, pageSize : number) :void{
    this.TopListItemsService.getTopListItems(pageId, pageSize).subscribe(
      (data) => {
        this.TopListItems =  data.topListItems;
        console.log(JSON.stringify(data))
      },
      (error) =>{
        console.error('Error fetching TopListItemsService:', error);
      }
    )
  };

  //HotTabsService
  fetchHotTabs(pageId: number, pageSize : number) :void{
    this.HotTabsService.getHotTabs(pageId, pageSize).subscribe(
      (data) => {
        this.hotListItems =  data;
        console.log(JSON.stringify(data))
      },
      (error) =>{
        console.error('Error fetching HotLists:', error);
      }
    )
  };

}
