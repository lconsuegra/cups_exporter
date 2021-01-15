package pkg

import (
	"github.com/phin1x/go-ipp"
	"github.com/prometheus/client_golang/prometheus"
)

func (e *Exporter) jobsMetrics(ch chan<- prometheus.Metric) error {

	printers, err := e.client.GetPrinters([]string{})

	if err != nil {
		e.log.Error(err, "failed to fetch printers")
		return err
	}

	for _, attr := range printers {

		if len(attr["printer-name"]) == 1 {

			printer := attr["printer-name"][0].Value.(string)

			jobs, err := e.client.GetJobs(printer, "", ipp.JobStateFilterAll, false, 0, 0, []string{})
			if err != nil {
				e.log.Error(err, "failed to fetch all jobs")
				return err
			}

			ch <- prometheus.MustNewConstMetric(e.jobsTotal, prometheus.CounterValue, float64(len(jobs)), printer)
		} else {
				e.log.Info("printer name attribute missing")
		}
	}

	return nil
}
