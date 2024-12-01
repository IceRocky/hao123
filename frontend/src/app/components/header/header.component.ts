import { Component } from '@angular/core';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.css']
})
export class HeaderComponent {
  dropdownVisible = false;
  searchDropdownVisible = false;

  searchInput: string = '';

  onSearchInputChange(event: Event): void {
    const inputElement = event.target as HTMLInputElement;
    this.searchInput = inputElement.value;
  }

  constructor() {}
  onSearch(): void {
    console.log('onSearch');
    window.open(`https://www.baidu.com/s?wd=${encodeURIComponent(this.searchInput)}`, '_blank');
  }


  toggleDropdown(): void {
    console.log('toggleDropdown');

    this.dropdownVisible = !this.dropdownVisible;
  }
  toggleSearchDropdown(): void {
    console.log('toggleSearchDropdown');
    this.searchDropdownVisible = !this.searchDropdownVisible;
  }
}
