import {Component} from '@angular/core'

interface Question {
  n1: number
  n2: number
  prod: number
  answer: number
}

type QuestionSet = Question[]

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {

  title = 'tables'
  tables = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12]
  currentTable = -1

  currentQuestionSet: QuestionSet = []
  currentQuestionIndex = 0
  currentQuestion: Question | null = null

  op1 = 1
  q: Question = {n1: 0, n2: 0, prod: 0, answer: 0}

  nextQuestion(): boolean {
    if (this.currentQuestionSet[this.currentQuestionIndex]) {
      this.currentQuestion = this.currentQuestionSet[this.currentQuestionIndex]
      this.currentQuestionIndex += 1
      return true
    }
    this.currentQuestionIndex = 0
    this.currentQuestion = null
    return false
  }

  practice(tablesNum: number): void {
    this.currentTable = tablesNum
    this._buildSet(tablesNum)
    // this.q = this.equation(this.op1, this.currentTable)
  }

  private _buildSet(tablesNum: number): void {
    this.currentQuestionSet = []
    this.tables.forEach((n, index) => {
      const q: Question = {
        n1: tablesNum,
        n2: n,
        prod: tablesNum * n,
        answer: 0
      }
      this.currentQuestionSet.push(q)
    })
  }

}
