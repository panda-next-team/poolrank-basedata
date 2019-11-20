<?php
namespace App\Observers;

class ModelObserver
{
    public function updating($model) {
        $model->updated_at_ts = time();
    }

    public function creating($model) {
        $model->created_at_ts = time();
        $model->updated_at_ts = 0;
    }
}
