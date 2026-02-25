import { Component, inject, OnInit } from '@angular/core';
import { Productservice } from '../services/productservice';
import { BaseChartDirective } from 'ng2-charts';
import { ChartConfiguration, ChartData, ChartType, registerables, Chart } from 'chart.js';
import { CommonModule } from '@angular/common';

Chart.register(...registerables);

interface SalesData {
  date: string;
  amount: number;
}

@Component({
  selector: 'app-saleschart',
  standalone: true,
  imports: [BaseChartDirective, CommonModule], 
  templateUrl: './saleschart.html',
  styleUrl: './saleschart.scss',
})

export class Saleschart implements OnInit {
  private productsService = inject(Productservice);
  
  public barChartType: ChartType = 'bar';
  public message: string = '';
  public today: Date = new Date();

  public barChartOptions: ChartConfiguration['options'] = {
    responsive: true,
    maintainAspectRatio: false,
    scales: {
      x: {
        ticks: { color: 'white' }, // X-axis labels
        grid: { color: 'rgba(255, 255, 255, 0.1)' } // Optional: light grid lines
      },
      y: {
        ticks: { color: 'white' }, // Y-axis labels
        grid: { color: 'rgba(255, 255, 255, 0.1)' }
      }
    },
    plugins: {
      legend: {
        labels: {
          color: 'white' // Legend text color (e.g., "Sales")
        }
      }
    }
  };

  public barChartData: ChartData<'bar'> = {
    labels: [],
    datasets: []
  };

  ngOnInit() {
    this.onViewReport();
  }

  onViewReport() {
    this.productsService.showSalesGraph().subscribe({
      next: (res: any) => {

        if (res.data?.errors) {
          this.message = res.data.errors[0].message;
          setTimeout(() => { this.message = ''; }, 3000);
          return;
        }

        const result = res.data?.salesChart;
        if (result) {
          this.barChartData = {
            labels: result.map((item: SalesData) => 
              new Date(item.date).toLocaleString('en-US', { month: 'short' })
            ),
            datasets: [{
              label: 'Sales',
              data: result.map((item: SalesData) => item.amount),
              backgroundColor: 'rgba(60, 179, 113)',
            }],
          };
        }
      },
      error: (err: any) => {
        this.message = err.message;
      }
    });
  }

  handlePrint() {
    window.print();
  }
}