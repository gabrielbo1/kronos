import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { KronosComponent } from './kronos.component';
import { PontoComponent } from './ponto/ponto.component';
import { AtendimentoComponent } from './atendimento/atendimento.component';
import { RelatorioComponent } from './relatorio/relatorio.component';
import { ConfiguracaoComponent } from './configuracao/configuracao.component';


const routes: Routes = [{
  path: '',
  component: KronosComponent,
  children: [
    {
      path: 'ponto',
      component: PontoComponent
    },
    {
      path: 'atendimento',
      component: AtendimentoComponent
    },
    {
      path: 'relatorio',
      component: RelatorioComponent
    },
    {
      path: 'configuracao',
      component: ConfiguracaoComponent
    }
  ]
}];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class KronosRoutingModule { }
