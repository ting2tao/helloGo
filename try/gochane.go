package main

func main() {
	//var wg sync.WaitGroup
	//mOut := make(chan struct{}, goNum) // 开goNum个并发
	//
	//for _i := range m {
	//	mOut <- struct{}{}
	//	wg.Add(1)
	//	go func(i *HiDeliveryRecord) {
	//		defer func() {
	//			wg.Done()
	//			<-mOut
	//		}()
	//		ctx.GetLogger().Info(ctx, zap.String("正在处理房屋code", i.oriHouses.Code), zap.String("正在处理房屋name", i.oriHouses.Name))
	//		if e, err := handleData(ctx, i); err != nil {
	//			/ctx.GetLogger().Error(ctx, zap.Error(err), zap.String("houseCode", i.oriHouses.Code))
	//			//atomic.AddInt64(&failHouse, 1)
	//			//failCodes = append(failCodes, i.oriHouses.Code)
	//		} else {
	//			//ctx.GetLogger().Info(ctx, zap.String("成功处理", e.Code()))
	//			/atomic.AddInt64(&successHouse, 1)
	//		}
	//	}(_i)
	//}
	//wg.Wait()
}
