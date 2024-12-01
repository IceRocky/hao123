import { Component } from '@angular/core';

@Component({
  selector: 'app-sideslide',
  templateUrl: './sideslide.component.html',
  styleUrls: ['./sideslide.component.css']
})
export class SideslideComponent {
  currentSlide = 0;
  totalSlides = 3; // Number of slides



  nextSlide() {
    this.currentSlide = (this.currentSlide + 1) % this.totalSlides;
  }

  prevSlide() {
    this.currentSlide =
      (this.currentSlide - 1 + this.totalSlides) % this.totalSlides;
  }
}
