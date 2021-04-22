import {ChangeDetectionStrategy, ChangeDetectorRef, Component, OnDestroy, OnInit} from '@angular/core';
import {map, takeUntil} from 'rxjs/operators';
import {Subject} from 'rxjs';
import {ActivatedRoute} from '@angular/router';
import {Location} from '@angular/common';
import {Root} from '../_models/root';
import {Trace} from '../_models/trace';
import {ApiService} from '../_services/api.service';
import {EventTypes} from '../_models/event-types';

@Component({
  selector: 'app-project-board',
  templateUrl: './evaluation-board.component.html',
  styleUrls: ['./evaluation-board.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush
})
export class EvaluationBoardComponent implements OnInit, OnDestroy {

  private unsubscribe$ = new Subject();

  public error: string = null;
  public contextId: string;
  public root: Root;
  public evaluations: Trace[];
  public hasHistory: boolean;

  constructor(private _changeDetectorRef: ChangeDetectorRef, private location: Location, private route: ActivatedRoute, private apiService: ApiService) {
    this.hasHistory = window.history.length > 1;
  }

  ngOnInit() {
    this.route.params
      .pipe(takeUntil(this.unsubscribe$))
      .subscribe(params => {
        if (params.shkeptncontext) {
          this.contextId = params.shkeptncontext;
          this.apiService.getTraces(this.contextId)
            .pipe(
              map(response => response.body),
              map(result => result.events || []),
              map(traces => traces.map(trace => Trace.fromJSON(trace)).sort((a, b) => new Date(a.time).getTime() - new Date(b.time).getTime()))
            )
            .pipe(takeUntil(this.unsubscribe$))
            .subscribe((traces: Trace[]) => {
              if (traces.length > 0) {
                this.root = Root.fromJSON(traces[0]);
                this.root.traces = traces;
                this.evaluations = traces.filter(t => t.type === EventTypes.EVALUATION_FINISHED && (!params.eventselector || t.id === params.eventselector || t.data.stage === params.eventselector)) ;
              } else {
                this.error = 'contextError';
              }
              this._changeDetectorRef.markForCheck();
            }, () => {
              this.error = 'error';
              this._changeDetectorRef.markForCheck();
            });
        }
      });
  }

  goBack(): void {
    this.location.back();
  }

  ngOnDestroy(): void {
    this.unsubscribe$.next();
    this.unsubscribe$.complete();
  }

}
