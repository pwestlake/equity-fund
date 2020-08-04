import { Component, ElementRef, Input, OnChanges, ViewChild, ViewEncapsulation, Output, EventEmitter } from '@angular/core';
import { DateValueModel } from 'src/data/date-value.model';
import * as d3 from 'd3';

@Component({
  selector: 'app-date-value',
  encapsulation: ViewEncapsulation.None,
  templateUrl: './date-value.component.html',
  styleUrls: ['./date-value.component.scss']
})
export class DateValueComponent implements OnChanges {
  @ViewChild('chart')
  private chartContainer: ElementRef;

  @Input()
  data: DateValueModel[];
  average: [number, number][];

  @Output() 
  viewMaximized = new EventEmitter<boolean>();

  @Output() 
  viewMinimized = new EventEmitter<boolean>();

  maxY: number = undefined;

  margin = {top: 20, right: 20, bottom: 30, left: 40};

  private y: number;

  constructor() { }

  ngOnChanges(): void {
    if (!this.data) { return; }

    this.createChart();
  }

  refresh() {
    setTimeout(() => {
      this.createChart();
    }, 0);
    
  }
  
  isMaxYUndefined(): boolean {
    return typeof(this.maxY) == 'undefined';
  }

  private createChart(): void {
    this.create7DayAverage();

    const element:HTMLElement = this.chartContainer.nativeElement;
    let child = element.lastElementChild;  
    while (child) { 
      element.removeChild(child); 
      child = element.lastElementChild; 
    } 
    const data = this.data;

    const svg = d3.select(element).append('svg')
        .attr('width', element.offsetWidth)
        .attr('height', element.offsetHeight);

    const contentWidth = element.offsetWidth - this.margin.left - this.margin.right;
    const contentHeight = element.offsetHeight - this.margin.top - this.margin.bottom;

    let startDate = new Date(data[0].date);
    let endDate = new Date(data[data.length - 1].date);

    const xLine = d3.scaleUtc()
      .domain(d3.extent(this.average, d => d[0]))
      .range([0, contentWidth])

    const yLine = d3.scaleLinear()
      .domain([0, d3.max(this.data, d => {
        if (typeof(this.maxY) == 'undefined') {
          return d.value
        }
        return d.value > this.maxY ? this.maxY : d.value;
        })])
      .range([contentHeight, this.margin.top])

    const line = d3.line()
      .x(d => xLine(d[0]))
      .y(d => yLine(d[1]))

    const x = d3
      .scaleBand()
      .range([0, contentWidth])
      .paddingInner(0.1)
      .align(0)
      .domain(data.map((d, i) => i.toString()));

    const xAxis = d3
      .scaleTime()
      .domain([startDate, endDate])
      .range([0, (x.step() * (data.length - 1)) ]);
  
    const y = d3
      .scaleLinear()
      .rangeRound([contentHeight, 0])
      .domain([0, d3.max(data, d => {
          if (typeof(this.maxY) == 'undefined') {
            return d.value
          }
          return d.value > this.maxY ? this.maxY : d.value;
        })]);

    let tooltip = d3.select(element).append("div")	
      .attr("class", "tooltip")				
      .style("opacity", 0);

    const g = svg.append('g')
      .attr('transform', `translate(${this.margin.left}, ${this.margin.top})`);
  
    g.append('g')
      .attr('class', 'axis axis--y')
      .call(d3.axisLeft(y).ticks(10))
      .append('text')
        .attr('transform', 'rotate(-90)')
        .attr('y', 6)
        .attr('dy', '0.71em')
        .attr('text-anchor', 'end')
        .text('Count');

    g.selectAll('.chart-primary')
      .data(data)
      .enter().append('rect')
        .attr('class', 'chart-primary')
        .attr('x', (d, i) => x(i.toString()))
        .attr('y', d => {
            if (typeof(this.maxY) == 'undefined') {
              return y(d.value);
            }
            return y(d.value > this.maxY ? this.maxY : d.value)
          })
        .attr('width', x.bandwidth)
        .attr('height', d => {
            if (typeof(this.maxY) == 'undefined') {
              return contentHeight - y(d.value);
            }
            return contentHeight - y(d.value > this.maxY ? this.maxY : d.value)
          })
        .attr('transform', 'translate(0,0)')
        .on("mouseover", d => this.showTooltip(d, tooltip, svg, element.offsetWidth))					
        .on("mouseout", d => this.hideTooltip(tooltip));

    g.append("path")
      .datum(this.average)
      .attr("fill", "none")
      .attr("stroke", "darkgrey")
      .attr("stroke-width", 2)
      .attr("stroke-linejoin", "round")
      .attr("stroke-linecap", "round")
      .attr("d", line);
    
      g.append('g')
      .attr('class', 'axis axis--x')
      .attr('transform', `translate(${x.bandwidth() / 2}, ${contentHeight})`)
      .call(d3.axisBottom(xAxis).ticks(5).tickFormat(d3.timeFormat("%b %d")));

      let setHeightBarFn = () : void => {return this.setHeightBar(this);};
      g.selectAll('.axis--x')
      .attr('fill', 'transparent')
      .on("click", setHeightBarFn);

      g.selectAll('.axis--y')
      .attr('fill', 'transparent')
      .on("click", function(d) {  console.log(d);})
  }

  zoomIn() {
    this.createChart();
  }

  zoomOut() {
    this.maxY = undefined;
    this.createChart();
  }

  maximized() {
    this.viewMaximized.emit(true);
  }

  minimized() {
    this.viewMinimized.emit(true);
  }

  setHeightBar(component: DateValueComponent) {
    const element:HTMLElement = component.chartContainer.nativeElement;
    const contentWidth = element.offsetWidth - component.margin.left - component.margin.right;
    const contentHeight = element.offsetHeight - component.margin.top - component.margin.bottom;

    const x = d3
      .scaleBand()
      .range([0, contentWidth])
      .paddingInner(0.1)
      .align(0)
      .domain(component.data.map((d, i) => i.toString()));

    const y = d3
      .scaleLinear()
      .rangeRound([contentHeight, 0])
      .domain([0, d3.max(component.data, d => d.value)]);
    
    
    const svgElement = element.getElementsByTagName("svg")[0];
    const svg = d3.select(svgElement);
    const context = d3.path();
    const maxValue = y(d3.max(component.data, d => d.value));
    const minValue = y(0);
    this.y = d3.mouse(svgElement)[1] - component.margin.top;
    context.moveTo(x("0"), this.y);
    context.lineTo(x((component.data.length - 1).toString()), this.y);
    
    let dragFn= () : void => {return component.dragged(svgElement, maxValue, minValue)};
    let dragEndFn = () : void => {return component.dragEnd(component, minValue, y)};
    svg.append("path")
      .attr("class", "v-measure")
      .attr("fill", "none")
      .attr("stroke", "darkgrey")
      .attr("stroke-width", 2)
      .attr("stroke-linejoin", "round")
      .attr("stroke-linecap", "round")
      .attr("d", context.toString())
      .attr('transform', `translate(${component.margin.left}, ${component.margin.top})`)
      .call(d3.drag()
        .on("drag", dragFn)
        .on("end", dragEndFn));
  }

  private dragged(d: SVGSVGElement, constraintMaxY: number, constraintMinY: number) {
    const vMeasure = d.getElementsByClassName("v-measure")[0];
    const selection = d3.select(vMeasure);

    let dy = d3.event.y - this.y;
    let coordY =  d3.mouse(d)[1];
    // if (coordY < this.margin.top) {
    //   coordY = this.margin.top;
    // }

    // if (coordY > constraintMinY + this.margin.top) {
    //   coordY = constraintMinY + this.margin.top;
    // }
    
    selection.attr('transform', `translate(${this.margin.left}, ${dy})`);
  }

  private dragEnd(component: DateValueComponent, constraintMinY: number, scaleY: any) {
    let y = d3.event.y;
     
    if (y < this.margin.top) {
      y = this.margin.top;
    }

    if (y > constraintMinY + this.margin.top) {
      y = constraintMinY + this.margin.top;
    }

    component.maxY = scaleY.invert(y - this.margin.top);
  }

  private showTooltip(d:any, tooltip:any, svg:any, maxWidth:any) {
    let coords = d3.mouse(svg.node());	
    let x = (coords[0] + 120) > maxWidth ? maxWidth - 120 : coords[0];

    tooltip.transition()		
        .duration(200)		
        .style("opacity", .9);		
    tooltip.html("<p>" + (new Date(d.date)).toDateString() + "</p><p>" + d.value + "</p>")	
        .style("left", x + "px")		
        .style("top", (coords[1]) + "px");	
  }

  private hideTooltip(tooltip:any) {
    tooltip.transition()		
    .duration(500)		
    .style("opacity", 0);	
  }

  private create7DayAverage() {
    this.average = new Array(this.data.length);
    
    let rollingAverage: number = 0;
    let value: number = 0;
    for (let i = 0; i < this.data.length; i++) {
      rollingAverage += (this.data[i].value / 7);

      if (i >= 7) {
        rollingAverage -= (this.data[i - 7].value / 7);
        value = rollingAverage;
      }

      const item: [number, number] = [new Date(this.data[i].date).getTime(), value];

      this.average[i] = item;
    }
  }
}
