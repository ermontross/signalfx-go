/*
 * Dashboards API
 *
 * API for creating, retrieving, updating, and deleting dashboards and dashboard groups. <br> Dashboards are groups of charts. In a dashboard, all the charts that belong to the dashboard appear at the same time and follow the same filtering options. ## Dashboard layout The system lays out dashboards and the charts they contain with these dimensions: <br>   * The web UI reserves a 12x100 grid for each dashboard and assigns one or     more charts to specific locations within the grid.   * A chart associated with the dashboard can be any size from 1x1 to 12x3.   * If you assign overlapping dashboard locations for charts, the system     attempts to resize or reorganize the layout. This ensures that all     of the charts fit within the space alloted to the dashboard.  ## Dashboard access By default, all users in an organization can edit and delete dashboards and dashboard groups. If SignalFx has enabled the \"write permissions\" feature for your organization, you can limit editing or deleting of specific dashboards to specific individuals or teams, or both. Use this feature to prevent unauthorized or accidental modifications to dashboards and the charts they contain. ## Cloning dashboards Users who don't have permission to edit a dashboard can still clone it and modify the clone. ## View dashboards You can view dashboards you create using the API in the web UI by specifying their \"id\" property in a web UI URL, by following this syntax: <br> <code>https://app.signalfx.com/#/dashboard/&lt;DASHBOARD_ID&gt;</code> <br> Dashboards you create using the API also appear by name in the web UI catalog and in their dashboard group.
 *
 * API version: 3.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dashboard

type OptionsHistogramChartOptions struct {
	// Color value to use as the theme of a histogram chart. The value is the index of the color in the standard SignalFx color palette, as shown in the following table. SignalFx displays all the datapoints using this color, but the chart shows smaller values with less color opacity and larger values with more opacity.<br> <table> <thead> <th style=\"text-align:center;\">Index</th><th>RGB hex value</th> </thead> <tbody> <tr><td>0</td><td>#999999</td></tr> <tr><td>1</td><td>#0077c2</td></tr> <tr><td>2</td><td>#00b9ff</td></tr> <tr><td>3</td><td>#6ca2b7</td></tr> <tr><td>4</td><td>#b04600</td></tr> <tr><td>5</td><td>#f47e00</td></tr> <tr><td>6</td><td>#e5b312</td></tr> <tr><td>7</td><td>#bd468d</td></tr> <tr><td>8</td><td>#e9008a</td></tr> <tr><td>9</td><td>#ff8dd1</td></tr> <tr><td>10</td><td>#876ff3</td></tr> <tr><td>11</td><td>#a747ff</td></tr> <tr><td>12</td><td>#ab99bc</td></tr> <tr><td>13</td><td>#007c1d</td></tr> <tr><td>14</td><td>#05ce00</td></tr> <tr><td>15</td><td>#0dba8f</td></tr> <tr><td>16</td><td>#ea1849</td></tr> <tr><td>17</td><td>#eac24b</td></tr> <tr><td>18</td><td>#e5e517</td></tr> <tr><td>19</td><td>#acef7f</td></tr> <tr><td>20</td><td>#6bd37e</td></tr> </tbody> </table> **Notes**   * Available if `options.type property` is `TimeSeriesChart`   * Available if `options.defaultPlotType` is `Histogram`   * Users may see colors other than those shown in the table, depending on the settings they select for color     blindness. To see sample swatches of the alternate colors and the mappings used for     color-blind users, see the color palette documentation at the end of the     [Charts Overview](https://developers.signalfx.com/reference#charts-overview-1).
	ColorThemeIndex int32 `json:"colorThemeIndex,omitempty"`
}